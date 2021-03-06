// Code generated by counterfeiter. DO NOT EDIT.
package ref

import (
	"net/http"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
)

type fakeHandleGetOrHeader struct {
	HandleGetOrHeadStub        func(dataHandle model.DataHandle, httpResp http.ResponseWriter, httpReq *http.Request)
	handleGetOrHeadMutex       sync.RWMutex
	handleGetOrHeadArgsForCall []struct {
		dataHandle model.DataHandle
		httpResp   http.ResponseWriter
		httpReq    *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeHandleGetOrHeader) HandleGetOrHead(dataHandle model.DataHandle, httpResp http.ResponseWriter, httpReq *http.Request) {
	fake.handleGetOrHeadMutex.Lock()
	fake.handleGetOrHeadArgsForCall = append(fake.handleGetOrHeadArgsForCall, struct {
		dataHandle model.DataHandle
		httpResp   http.ResponseWriter
		httpReq    *http.Request
	}{dataHandle, httpResp, httpReq})
	fake.recordInvocation("HandleGetOrHead", []interface{}{dataHandle, httpResp, httpReq})
	fake.handleGetOrHeadMutex.Unlock()
	if fake.HandleGetOrHeadStub != nil {
		fake.HandleGetOrHeadStub(dataHandle, httpResp, httpReq)
	}
}

func (fake *fakeHandleGetOrHeader) HandleGetOrHeadCallCount() int {
	fake.handleGetOrHeadMutex.RLock()
	defer fake.handleGetOrHeadMutex.RUnlock()
	return len(fake.handleGetOrHeadArgsForCall)
}

func (fake *fakeHandleGetOrHeader) HandleGetOrHeadArgsForCall(i int) (model.DataHandle, http.ResponseWriter, *http.Request) {
	fake.handleGetOrHeadMutex.RLock()
	defer fake.handleGetOrHeadMutex.RUnlock()
	return fake.handleGetOrHeadArgsForCall[i].dataHandle, fake.handleGetOrHeadArgsForCall[i].httpResp, fake.handleGetOrHeadArgsForCall[i].httpReq
}

func (fake *fakeHandleGetOrHeader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleGetOrHeadMutex.RLock()
	defer fake.handleGetOrHeadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeHandleGetOrHeader) recordInvocation(key string, args []interface{}) {
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
