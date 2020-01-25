// Code generated by counterfeiter. DO NOT EDIT.
package data

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
)

type FakeProvider struct {
	TryResolveStub        func(context.Context, string) (model.DataHandle, error)
	tryResolveMutex       sync.RWMutex
	tryResolveArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	tryResolveReturns struct {
		result1 model.DataHandle
		result2 error
	}
	tryResolveReturnsOnCall map[int]struct {
		result1 model.DataHandle
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProvider) TryResolve(arg1 context.Context, arg2 string) (model.DataHandle, error) {
	fake.tryResolveMutex.Lock()
	ret, specificReturn := fake.tryResolveReturnsOnCall[len(fake.tryResolveArgsForCall)]
	fake.tryResolveArgsForCall = append(fake.tryResolveArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("TryResolve", []interface{}{arg1, arg2})
	fake.tryResolveMutex.Unlock()
	if fake.TryResolveStub != nil {
		return fake.TryResolveStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.tryResolveReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProvider) TryResolveCallCount() int {
	fake.tryResolveMutex.RLock()
	defer fake.tryResolveMutex.RUnlock()
	return len(fake.tryResolveArgsForCall)
}

func (fake *FakeProvider) TryResolveCalls(stub func(context.Context, string) (model.DataHandle, error)) {
	fake.tryResolveMutex.Lock()
	defer fake.tryResolveMutex.Unlock()
	fake.TryResolveStub = stub
}

func (fake *FakeProvider) TryResolveArgsForCall(i int) (context.Context, string) {
	fake.tryResolveMutex.RLock()
	defer fake.tryResolveMutex.RUnlock()
	argsForCall := fake.tryResolveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProvider) TryResolveReturns(result1 model.DataHandle, result2 error) {
	fake.tryResolveMutex.Lock()
	defer fake.tryResolveMutex.Unlock()
	fake.TryResolveStub = nil
	fake.tryResolveReturns = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeProvider) TryResolveReturnsOnCall(i int, result1 model.DataHandle, result2 error) {
	fake.tryResolveMutex.Lock()
	defer fake.tryResolveMutex.Unlock()
	fake.TryResolveStub = nil
	if fake.tryResolveReturnsOnCall == nil {
		fake.tryResolveReturnsOnCall = make(map[int]struct {
			result1 model.DataHandle
			result2 error
		})
	}
	fake.tryResolveReturnsOnCall[i] = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tryResolveMutex.RLock()
	defer fake.tryResolveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProvider) recordInvocation(key string, args []interface{}) {
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

var _ Provider = new(FakeProvider)
