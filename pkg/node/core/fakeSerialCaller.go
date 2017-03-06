// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/pkg/model"
)

type fakeSerialCaller struct {
	CallStub        func(inboundScope map[string]*model.Data, rootOpId string, opPkgRef string, scgSerialCall []*model.Scg) (outboundScope map[string]*model.Data, err error)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		inboundScope  map[string]*model.Data
		rootOpId      string
		opPkgRef      string
		scgSerialCall []*model.Scg
	}
	callReturns struct {
		result1 map[string]*model.Data
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeSerialCaller) Call(inboundScope map[string]*model.Data, rootOpId string, opPkgRef string, scgSerialCall []*model.Scg) (outboundScope map[string]*model.Data, err error) {
	var scgSerialCallCopy []*model.Scg
	if scgSerialCall != nil {
		scgSerialCallCopy = make([]*model.Scg, len(scgSerialCall))
		copy(scgSerialCallCopy, scgSerialCall)
	}
	fake.callMutex.Lock()
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		inboundScope  map[string]*model.Data
		rootOpId      string
		opPkgRef      string
		scgSerialCall []*model.Scg
	}{inboundScope, rootOpId, opPkgRef, scgSerialCallCopy})
	fake.recordInvocation("Call", []interface{}{inboundScope, rootOpId, opPkgRef, scgSerialCallCopy})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(inboundScope, rootOpId, opPkgRef, scgSerialCall)
	}
	return fake.callReturns.result1, fake.callReturns.result2
}

func (fake *fakeSerialCaller) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *fakeSerialCaller) CallArgsForCall(i int) (map[string]*model.Data, string, string, []*model.Scg) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.callArgsForCall[i].inboundScope, fake.callArgsForCall[i].rootOpId, fake.callArgsForCall[i].opPkgRef, fake.callArgsForCall[i].scgSerialCall
}

func (fake *fakeSerialCaller) CallReturns(result1 map[string]*model.Data, result2 error) {
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 map[string]*model.Data
		result2 error
	}{result1, result2}
}

func (fake *fakeSerialCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeSerialCaller) recordInvocation(key string, args []interface{}) {
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
