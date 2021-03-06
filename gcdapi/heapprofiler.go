// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains HeapProfiler functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/zurrix/gcd/gcdmessage"
)

// Sampling Heap Profile node. Holds callsite information, allocation statistics and child nodes.
type HeapProfilerSamplingHeapProfileNode struct {
	FunctionName string                                 `json:"functionName"` // Function name.
	ScriptId     string                                 `json:"scriptId"`     // Script identifier.
	Url          string                                 `json:"url"`          // URL.
	LineNumber   int                                    `json:"lineNumber"`   // 1-based line number of the function start position.
	ColumnNumber int                                    `json:"columnNumber"` // 1-based column number of the function start position.
	SelfSize     float64                                `json:"selfSize"`     // Allocations size in bytes for the node excluding children.
	Children     []*HeapProfilerSamplingHeapProfileNode `json:"children"`     // Child nodes.
}

// Profile.
type HeapProfilerSamplingHeapProfile struct {
	Head *HeapProfilerSamplingHeapProfileNode `json:"head"` //
}

//
type HeapProfilerAddHeapSnapshotChunkEvent struct {
	Method string `json:"method"`
	Params struct {
		Chunk string `json:"chunk"` //
	} `json:"Params,omitempty"`
}

//
type HeapProfilerReportHeapSnapshotProgressEvent struct {
	Method string `json:"method"`
	Params struct {
		Done     int  `json:"done"`               //
		Total    int  `json:"total"`              //
		Finished bool `json:"finished,omitempty"` //
	} `json:"Params,omitempty"`
}

// If heap objects tracking has been started then backend regulary sends a current value for last seen object id and corresponding timestamp. If the were changes in the heap since last event then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
type HeapProfilerLastSeenObjectIdEvent struct {
	Method string `json:"method"`
	Params struct {
		LastSeenObjectId int     `json:"lastSeenObjectId"` //
		Timestamp        float64 `json:"timestamp"`        //
	} `json:"Params,omitempty"`
}

// If heap objects tracking has been started then backend may send update for one or more fragments
type HeapProfilerHeapStatsUpdateEvent struct {
	Method string `json:"method"`
	Params struct {
		StatsUpdate []int `json:"statsUpdate"` // An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
	} `json:"Params,omitempty"`
}

type HeapProfiler struct {
	target gcdmessage.ChromeTargeter
}

func NewHeapProfiler(target gcdmessage.ChromeTargeter) *HeapProfiler {
	c := &HeapProfiler{target: target}
	return c
}

//
func (c *HeapProfiler) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.enable"})
}

//
func (c *HeapProfiler) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.disable"})
}

// StartTrackingHeapObjects -
// trackAllocations -
func (c *HeapProfiler) StartTrackingHeapObjects(trackAllocations bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["trackAllocations"] = trackAllocations
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.startTrackingHeapObjects", Params: paramRequest})
}

// StopTrackingHeapObjects -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
func (c *HeapProfiler) StopTrackingHeapObjects(reportProgress bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.stopTrackingHeapObjects", Params: paramRequest})
}

// TakeHeapSnapshot -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
func (c *HeapProfiler) TakeHeapSnapshot(reportProgress bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.takeHeapSnapshot", Params: paramRequest})
}

//
func (c *HeapProfiler) CollectGarbage() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.collectGarbage"})
}

// GetObjectByHeapObjectId -
// objectId -
// objectGroup - Symbolic group name that can be used to release multiple objects.
// Returns -  result - Evaluation result.
func (c *HeapProfiler) GetObjectByHeapObjectId(objectId string, objectGroup string) (*RuntimeRemoteObject, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["objectId"] = objectId
	paramRequest["objectGroup"] = objectGroup
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getObjectByHeapObjectId", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result *RuntimeRemoteObject
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Result, nil
}

// AddInspectedHeapObject - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
// heapObjectId - Heap snapshot object id to be accessible by means of $x command line API.
func (c *HeapProfiler) AddInspectedHeapObject(heapObjectId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["heapObjectId"] = heapObjectId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.addInspectedHeapObject", Params: paramRequest})
}

// GetHeapObjectId -
// objectId - Identifier of the object to get heap object id for.
// Returns -  heapSnapshotObjectId - Id of the heap snapshot object corresponding to the passed remote object id.
func (c *HeapProfiler) GetHeapObjectId(objectId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getHeapObjectId", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			HeapSnapshotObjectId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.HeapSnapshotObjectId, nil
}

// StartSampling -
// samplingInterval - Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
func (c *HeapProfiler) StartSampling(samplingInterval float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["samplingInterval"] = samplingInterval
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.startSampling", Params: paramRequest})
}

// StopSampling -
// Returns -  profile - Recorded sampling heap profile.
func (c *HeapProfiler) StopSampling() (*HeapProfilerSamplingHeapProfile, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.stopSampling"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Profile *HeapProfilerSamplingHeapProfile
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Profile, nil
}
