package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	pb "github.com/hotnops/gtunnel/gTunnel"
)

type TaskStream interface {
	Send(*pb.EndpointTaskMessage) error
	Recv() (*pb.EndpointTaskMessage, error)
}

type TaskExecutor interface {
	Execute(*Task, *TaskHandler) error
}

// type TaskExecutor func(*Task) TaskResponse

type Task struct {
	*pb.EndpointTaskMessage
	AuxData   map[uint32][]byte
	Executing chan int
}

func (t *Task) String() (out string) {
	if t.TaskType == TaskCMD {
		out = fmt.Sprintf("%s", t.Cmd)
	} else if t.IsFileTransfer() {
		out = fmt.Sprintf("file=%s currentChunk=%d status=%d", t.Cmd, t.Chunk, t.Status)
	}
	return
}

// Whem gServer receives a result
func (t *Task) Process(taskMsg *pb.EndpointTaskMessage) {
	if t.Status == TaskStatusDone {
		if t.TaskType == TaskCMD {
			fmt.Printf("Task [%s] finished\n%s", t.TaskId, string(taskMsg.Data))
		} else if t.TaskType == TaskUploadFile {
			log.Print("\n[File uploaded successfuly]")
		} else if t.TaskType == TaskDownloadFile {
			// TODO save file from aux buffer
			// for chunk, _ := range t.AuxData {
			// 	fmt.Println(chunk)
			// }
		} else if t.TaskType == TaskTCPScan {
			log.Print(string(taskMsg.Data))
		}
	} else if t.Status == TaskStatusError {
		fmt.Printf("Task [%s] errored with: \n%s\n", t.TaskId, string(taskMsg.Data))
	} else if t.Status == TaskStatusRunning {
		if t.TaskType == TaskDownloadFile {
			t.AuxData[uint32(taskMsg.Chunk)] = taskMsg.Data
		} else if t.TaskType == TaskTCPScan {
			// log.Print(string(taskMsg.Data))
			fmt.Println(string(taskMsg.Data))
		}
	} else if t.Status == TaskStatusFileDone {
		if t.TaskType == TaskDownloadFile {
			file := filepath.Base(t.Cmd)
			f, err := os.Create(file)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()
			for i := 0; i < len(t.AuxData); i++ {
				_, err := f.Write(t.AuxData[uint32(i)])
				if err != nil {
					t.Status = TaskStatusError
					return
				}
			}
			t.Status = TaskStatusDone
			log.Print("\n[File downloaded successfuly]")
		}
	}
}

// A structure to handle the TCP connection
// and map them to the gRPC byte stream.
type TaskHandler struct {
	Kill         chan bool
	Status       int32
	endpointId   string
	ingressTasks chan *pb.EndpointTaskMessage
	taskStream   TaskStream
	bytesTx      uint64
	bytesRx      uint64
	Tasks        map[string]*Task
	executioner  bool
	taskExecutor TaskExecutor
	remoteClose  bool
}

func (t *Task) IsFileTransfer() bool {
	return t.TaskType == TaskDownloadFile || t.TaskType == TaskUploadFile
}

// SetStream will set the byteStream for a connection
func (t *TaskHandler) SetStream(s TaskStream) {
	t.taskStream = s
}

func (t *TaskHandler) Close() {
	if t.Status != ConnectionStatusClosed {
		close(t.Kill)
		t.Status = ConnectionStatusClosed
	}
}

func (t *TaskHandler) IsOwnMessage(msg *pb.EndpointTaskMessage) bool {
	return msg.EndpointID == t.endpointId
}

// handleIngressData will handle all incoming messages
// on the gRPC byte stream and send them to the locally
// connected socket.
func (t *TaskHandler) handleIngressData() {

	inputChan := make(chan *pb.EndpointTaskMessage)

	go func(s TaskStream) {
		for {
			message, err := s.Recv()
			if err != nil {
				t.Close()
				return
			}
			//
			if !t.IsOwnMessage(message) {
				inputChan <- message
			}
		}
	}(t.taskStream)

	for {
		select {
		case taskMessage, ok := <-inputChan:
			if !ok {
				inputChan = nil
				break
			}
			if t.IsMessageValid(taskMessage) {
				// When response is receives / most likely from gServer
				if task, ok := t.Tasks[taskMessage.TaskId]; ok {
					// First we sync the local Task with the new message
					task.EndpointTaskMessage = taskMessage
					task.Process(taskMessage)
					if task.Status == TaskStatusDone {
						delete(t.Tasks, task.TaskId)
					} else if (task.Status == TaskStatusRunning || task.Status == TaskStatusFileDone) && task.IsFileTransfer() && t.executioner {
						t.RunTask(task)
					}
				} else {
					// When request for a new Task is received
					// most likely from gClient
					go t.NewTask(taskMessage)
				}
			}
		}
		if inputChan == nil {
			fmt.Println("wat")
			break
		}
	}
}

func (t *TaskHandler) IsMessageValid(message *pb.EndpointTaskMessage) (isValid bool) {
	// TODO validate
	return true
}

func (t *TaskHandler) SetExecutor(taskExecutor TaskExecutor) {
	t.taskExecutor = taskExecutor
}

// RunTask executed within the gClient
func (t *TaskHandler) RunTask(task *Task) {
	if t.taskExecutor != nil {
		err := t.taskExecutor.Execute(task, t)
		if err != nil {
			fmt.Println(err)
		}
		if task.Status == TaskStatusDone || task.Status == TaskStatusError {
			// Publishes the result
			t.SendTaskMessage(task.EndpointTaskMessage)
			delete(t.Tasks, task.TaskId)
		}
	} else {
		fmt.Printf("Attempt to execute task %s but Executor is nil\n", task.TaskId)
	}
}

// NewTask runs within gClient
func (t *TaskHandler) NewTask(message *pb.EndpointTaskMessage) {
	if !t.executioner {
		return
	}
	fmt.Println("New Task received: ", message.TaskId)
	t.Tasks[message.TaskId] = &Task{
		EndpointTaskMessage: message,
		AuxData:             make(map[uint32][]byte),
	}

	t.RunTask(t.Tasks[message.TaskId])
	return
}

// NewTask runs within gServer
func (t *TaskHandler) SendTask(task *Task) {
	t.Tasks[task.TaskId] = task
	t.SendTaskMessage(task.EndpointTaskMessage)
}

// SendCloseMessage will send a zero sized
// message to the remote endpoint, indicating
// that a TCP connection has been closed locally.
func (t *TaskHandler) SendTaskMessage(task *pb.EndpointTaskMessage) {
	task.EndpointID = t.endpointId
	t.taskStream.Send(task)
}

// Start will start two goroutines for handling the TCP socket
// and the gRPC stream.
func (t *TaskHandler) Start() {
	go t.handleIngressData()
}
