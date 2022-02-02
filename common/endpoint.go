package common

import (
	pb "github.com/hotnops/gtunnel/gTunnel"
)

type TunnelControlStream interface {
	Send(*pb.TunnelControlMessage) error
	Recv() (*pb.TunnelControlMessage, error)
}

type Endpoint struct {
	Id                 string
	hostname           string
	killClient         chan bool
	tunnels            map[string]*Tunnel
	endpointCtrlStream chan pb.EndpointControlMessage
	TaskHandler        *TaskHandler
}

// AddTunnel adds a tunnel instance to the list of tunnels
// maintained by the endpoint
func (e *Endpoint) AddTunnel(id string, t *Tunnel) {
	e.tunnels[id] = t
}

// GetTunnels returns all of the active tunnels
// maintained by the endpoint
func (e *Endpoint) GetTunnels() map[string]*Tunnel {
	return e.tunnels
}

//GetTunnel will take in a tunnel ID string as an argument
// and return a Tunnel pointer of the corresponding ID.
func (e *Endpoint) GetTunnel(tunID string) (t *Tunnel, ok bool) {
	if e.tunnels == nil {
		return nil, false
	}
	if t, ok = e.tunnels[tunID]; !ok {
		return nil, ok
	}
	return
}

// RemoveTunnel takes in a tunnelID as an argument
// and removes the tunnel from the endpoint. Returns
// true if successful and false otherwise.
func (e *Endpoint) RemoveTunnel(tunID string) bool {
	tun, ok := e.tunnels[tunID]
	if !ok {
		return false
	}
	tun.Stop()
	delete(e.tunnels, tunID)
	return true
}

// Stop will close all tunnels and the associated TCP
// connections with each tunnel.
func (e *Endpoint) Stop() {
	for id, _ := range e.tunnels {
		e.RemoveTunnel(id)
	}
	close(e.endpointCtrlStream)
}

// NewTaskHandler is a constructor function for a new TaskHandler.
func (e *Endpoint) NewTaskHandler(s TaskStream, executioner bool, ownEndpointId string) (err error) {
	th := &TaskHandler{
		executioner: executioner,
		taskStream:  s,
		endpointId:  ownEndpointId,
		Tasks:       make(map[string]*Task),
		Kill:        make(chan bool),
	}
	e.TaskHandler = th
	return
}

// NewEndpoint is a constructor for the endpoint
// class. It takes an endpoint ID string as an argument.
func NewEndpoint(id string) *Endpoint {
	e := new(Endpoint)
	e.Id = id
	e.endpointCtrlStream = make(chan pb.EndpointControlMessage)
	e.tunnels = make(map[string]*Tunnel)
	return e
}
