// Code generated by counterfeiter. DO NOT EDIT.
package manifest

import (
	"sync"
)

type fakeValidator struct {
	ValidateStub        func(manifestBytes []byte) []error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		manifestBytes []byte
	}
	validateReturns struct {
		result1 []error
	}
	validateReturnsOnCall map[int]struct {
		result1 []error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeValidator) Validate(manifestBytes []byte) []error {
	var manifestBytesCopy []byte
	if manifestBytes != nil {
		manifestBytesCopy = make([]byte, len(manifestBytes))
		copy(manifestBytesCopy, manifestBytes)
	}
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		manifestBytes []byte
	}{manifestBytesCopy})
	fake.recordInvocation("Validate", []interface{}{manifestBytesCopy})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(manifestBytes)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateReturns.result1
}

func (fake *fakeValidator) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *fakeValidator) ValidateArgsForCall(i int) []byte {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].manifestBytes
}

func (fake *fakeValidator) ValidateReturns(result1 []error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 []error
	}{result1}
}

func (fake *fakeValidator) ValidateReturnsOnCall(i int, result1 []error) {
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 []error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 []error
	}{result1}
}

func (fake *fakeValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeValidator) recordInvocation(key string, args []interface{}) {
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

var _ Validator = new(fakeValidator)
