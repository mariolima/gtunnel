package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/hotnops/gtunnel/common"
	"github.com/zs5460/ipconv"
)

type Executor struct{}

func (e *Executor) executeCmd(command string) (result []byte, err error) {
	data := strings.Split(command, " ")
	cmd := exec.Command(data[0], data[1:]...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	result, err = cmd.CombinedOutput()
	return
}

func (e *Executor) Failed(task *common.Task, handler *common.TaskHandler, err error) error {
	fmt.Println(err)
	task.Status = common.TaskStatusError
	task.Data = []byte(err.Error())
	// handler.SendTask(task)
	return err
}

func (e *Executor) Execute(task *common.Task, handler *common.TaskHandler) (err error) {
	// task.Status = common.TaskStatusRunning
	if task.TaskType == common.TaskCMD {
		var result []byte
		result, err = e.executeCmd(task.Cmd)
		if err != nil {
			return e.Failed(task, handler, err)
		}
		task.Data = result
		task.Status = common.TaskStatusDone
	} else if task.TaskType == common.TaskUploadFile {
		if task.Status == common.TaskStatusFileDone {
			var f *os.File
			f, err = os.Create(task.Cmd)
			if err != nil {
				return
			}
			defer f.Close()
			for i := 0; i < len(task.AuxData); i++ {
				_, err = f.Write(task.AuxData[uint32(i)])
				if err != nil {
					task.Status = common.TaskStatusError
					return
				}
			}
			task.Status = common.TaskStatusDone
			task.Data = []byte("")
		} else {
			task.AuxData[uint32(task.Chunk)] = task.Data
			fmt.Printf("loaded %d chunk into auxData\n", task.Chunk)
		}
	} else if task.TaskType == common.TaskDownloadFile {
		var f *os.File
		f, err = os.Open(task.Cmd)
		if err != nil {
			return
		}
		buf := make([]byte, 3*1024)
		var (
			writing = true
			n       int
			chunk   int
		)
		for writing {
			n, err = f.Read(buf)
			if err != nil {
				if err == io.EOF {
					writing = false
					err = nil
					continue
				}
				return
			}
			task.Data = buf[:n]
			task.Chunk = int32(chunk)
			// fmt.Println(task.Chunk)
			handler.SendTask(task)
			chunk += 1
		}
		task.Data = []byte{}
		task.Status = common.TaskStatusFileDone
		handler.SendTask(task)
	} else if task.TaskType == common.TaskTCPScan {
		opts := new(common.TaskTCPScanOpts)
		fmt.Println(task.Data)
		err = json.Unmarshal(task.Data, opts)
		if err != nil {
			return e.Failed(task, handler, err)
		}

		runtime.GOMAXPROCS(runtime.NumCPU())
		maxThread := 10000
		fullMode := false
		if opts.Port == "all" {
			fullMode = true
		}

		specifiedPort := opts.Port
		fmt.Println("running scan " + opts.Ipnet)
		ips, err := ipconv.Parse(opts.Ipnet)
		if err != nil {
			return e.Failed(task, handler, err)
		}

		result := make(chan string, 1024)
		done := make(chan int)

		//to pad a new line in the cli
		task.Data = []byte("")
		handler.SendTask(task)
		go func() {
			for s := range result {
				task.Data = []byte(s)
				handler.SendTask(task)
				fmt.Println(s)
			}
		}()

		tasks := make(chan Item)
		go GenTask(specifiedPort, fullMode, tasks, ips)
		go Scan(maxThread, tasks, done, result)
		<-done
		task.Status = common.TaskStatusDone
		task.Data = []byte("Scan Ended")
		handler.SendTask(task)
	} else if task.TaskType == common.TaskRunSC {
		// TODO
	}
	return
}
