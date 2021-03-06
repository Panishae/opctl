// Code generated by counterfeiter. DO NOT EDIT.
package path

import (
	"net/http"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
)

type FakeHandler struct {
	HandleStub        func(dataHandle model.DataHandle, dataPath string, httpResp http.ResponseWriter, httpReq *http.Request)
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
		dataHandle model.DataHandle
		dataPath   string
		httpResp   http.ResponseWriter
		httpReq    *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandler) Handle(dataHandle model.DataHandle, dataPath string, httpResp http.ResponseWriter, httpReq *http.Request) {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
		dataHandle model.DataHandle
		dataPath   string
		httpResp   http.ResponseWriter
		httpReq    *http.Request
	}{dataHandle, dataPath, httpResp, httpReq})
	fake.recordInvocation("Handle", []interface{}{dataHandle, dataPath, httpResp, httpReq})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		fake.HandleStub(dataHandle, dataPath, httpResp, httpReq)
	}
}

func (fake *FakeHandler) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeHandler) HandleArgsForCall(i int) (model.DataHandle, string, http.ResponseWriter, *http.Request) {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return fake.handleArgsForCall[i].dataHandle, fake.handleArgsForCall[i].dataPath, fake.handleArgsForCall[i].httpResp, fake.handleArgsForCall[i].httpReq
}

func (fake *FakeHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ Handler = new(FakeHandler)
