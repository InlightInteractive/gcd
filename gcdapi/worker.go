// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Worker functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/zurrix/gcd/gcdmessage"
)

//
type WorkerWorkerCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		WorkerId           string `json:"workerId"`           //
		Url                string `json:"url"`                //
		WaitingForDebugger bool   `json:"waitingForDebugger"` //
	} `json:"Params,omitempty"`
}

//
type WorkerWorkerTerminatedEvent struct {
	Method string `json:"method"`
	Params struct {
		WorkerId string `json:"workerId"` //
	} `json:"Params,omitempty"`
}

//
type WorkerDispatchMessageFromWorkerEvent struct {
	Method string `json:"method"`
	Params struct {
		WorkerId string `json:"workerId"` //
		Message  string `json:"message"`  //
	} `json:"Params,omitempty"`
}

type Worker struct {
	target gcdmessage.ChromeTargeter
}

func NewWorker(target gcdmessage.ChromeTargeter) *Worker {
	c := &Worker{target: target}
	return c
}

//
func (c *Worker) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.enable"})
}

//
func (c *Worker) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.disable"})
}

// SendMessageToWorker -
// workerId -
// message -
func (c *Worker) SendMessageToWorker(workerId string, message string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["workerId"] = workerId
	paramRequest["message"] = message
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.sendMessageToWorker", Params: paramRequest})
}

// SetWaitForDebuggerOnStart -
// value -
func (c *Worker) SetWaitForDebuggerOnStart(value bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.setWaitForDebuggerOnStart", Params: paramRequest})
}
