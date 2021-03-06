// Code generated by counterfeiter. DO NOT EDIT.
package docker

import (
	"sync"
)

type fakeFSPathConverter struct {
	LocalToEngineStub        func(localPath string) string
	localToEngineMutex       sync.RWMutex
	localToEngineArgsForCall []struct {
		localPath string
	}
	localToEngineReturns struct {
		result1 string
	}
	localToEngineReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeFSPathConverter) LocalToEngine(localPath string) string {
	fake.localToEngineMutex.Lock()
	ret, specificReturn := fake.localToEngineReturnsOnCall[len(fake.localToEngineArgsForCall)]
	fake.localToEngineArgsForCall = append(fake.localToEngineArgsForCall, struct {
		localPath string
	}{localPath})
	fake.recordInvocation("LocalToEngine", []interface{}{localPath})
	fake.localToEngineMutex.Unlock()
	if fake.LocalToEngineStub != nil {
		return fake.LocalToEngineStub(localPath)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.localToEngineReturns.result1
}

func (fake *fakeFSPathConverter) LocalToEngineCallCount() int {
	fake.localToEngineMutex.RLock()
	defer fake.localToEngineMutex.RUnlock()
	return len(fake.localToEngineArgsForCall)
}

func (fake *fakeFSPathConverter) LocalToEngineArgsForCall(i int) string {
	fake.localToEngineMutex.RLock()
	defer fake.localToEngineMutex.RUnlock()
	return fake.localToEngineArgsForCall[i].localPath
}

func (fake *fakeFSPathConverter) LocalToEngineReturns(result1 string) {
	fake.LocalToEngineStub = nil
	fake.localToEngineReturns = struct {
		result1 string
	}{result1}
}

func (fake *fakeFSPathConverter) LocalToEngineReturnsOnCall(i int, result1 string) {
	fake.LocalToEngineStub = nil
	if fake.localToEngineReturnsOnCall == nil {
		fake.localToEngineReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.localToEngineReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *fakeFSPathConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.localToEngineMutex.RLock()
	defer fake.localToEngineMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeFSPathConverter) recordInvocation(key string, args []interface{}) {
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
