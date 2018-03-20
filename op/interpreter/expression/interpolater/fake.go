// Code generated by counterfeiter. DO NOT EDIT.
package interpolater

import (
	"github.com/opspec-io/sdk-golang/model"
	"sync"
)

type Fake struct {
	InterpolateStub        func(expression string, scope map[string]*model.Value, opHandle model.DataHandle) (string, error)
	interpolateMutex       sync.RWMutex
	interpolateArgsForCall []struct {
		expression string
		scope      map[string]*model.Value
		opHandle   model.DataHandle
	}
	interpolateReturns struct {
		result1 string
		result2 error
	}
	interpolateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Interpolate(expression string, scope map[string]*model.Value, opHandle model.DataHandle) (string, error) {
	fake.interpolateMutex.Lock()
	ret, specificReturn := fake.interpolateReturnsOnCall[len(fake.interpolateArgsForCall)]
	fake.interpolateArgsForCall = append(fake.interpolateArgsForCall, struct {
		expression string
		scope      map[string]*model.Value
		opHandle   model.DataHandle
	}{expression, scope, opHandle})
	fake.recordInvocation("Interpolate", []interface{}{expression, scope, opHandle})
	fake.interpolateMutex.Unlock()
	if fake.InterpolateStub != nil {
		return fake.InterpolateStub(expression, scope, opHandle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpolateReturns.result1, fake.interpolateReturns.result2
}

func (fake *Fake) InterpolateCallCount() int {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return len(fake.interpolateArgsForCall)
}

func (fake *Fake) InterpolateArgsForCall(i int) (string, map[string]*model.Value, model.DataHandle) {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return fake.interpolateArgsForCall[i].expression, fake.interpolateArgsForCall[i].scope, fake.interpolateArgsForCall[i].opHandle
}

func (fake *Fake) InterpolateReturns(result1 string, result2 error) {
	fake.InterpolateStub = nil
	fake.interpolateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) InterpolateReturnsOnCall(i int, result1 string, result2 error) {
	fake.InterpolateStub = nil
	if fake.interpolateReturnsOnCall == nil {
		fake.interpolateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.interpolateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ Interpolater = new(Fake)
