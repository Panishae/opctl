// Code generated by counterfeiter. DO NOT EDIT.
package expression

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakeEvalArrayInitializerer struct {
	EvalStub        func(expression []interface{}, scope map[string]*model.Value, opHandle model.DataHandle) ([]interface{}, error)
	evalMutex       sync.RWMutex
	evalArgsForCall []struct {
		expression []interface{}
		scope      map[string]*model.Value
		opHandle   model.DataHandle
	}
	evalReturns struct {
		result1 []interface{}
		result2 error
	}
	evalReturnsOnCall map[int]struct {
		result1 []interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeEvalArrayInitializerer) Eval(expression []interface{}, scope map[string]*model.Value, opHandle model.DataHandle) ([]interface{}, error) {
	var expressionCopy []interface{}
	if expression != nil {
		expressionCopy = make([]interface{}, len(expression))
		copy(expressionCopy, expression)
	}
	fake.evalMutex.Lock()
	ret, specificReturn := fake.evalReturnsOnCall[len(fake.evalArgsForCall)]
	fake.evalArgsForCall = append(fake.evalArgsForCall, struct {
		expression []interface{}
		scope      map[string]*model.Value
		opHandle   model.DataHandle
	}{expressionCopy, scope, opHandle})
	fake.recordInvocation("Eval", []interface{}{expressionCopy, scope, opHandle})
	fake.evalMutex.Unlock()
	if fake.EvalStub != nil {
		return fake.EvalStub(expression, scope, opHandle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.evalReturns.result1, fake.evalReturns.result2
}

func (fake *fakeEvalArrayInitializerer) EvalCallCount() int {
	fake.evalMutex.RLock()
	defer fake.evalMutex.RUnlock()
	return len(fake.evalArgsForCall)
}

func (fake *fakeEvalArrayInitializerer) EvalArgsForCall(i int) ([]interface{}, map[string]*model.Value, model.DataHandle) {
	fake.evalMutex.RLock()
	defer fake.evalMutex.RUnlock()
	return fake.evalArgsForCall[i].expression, fake.evalArgsForCall[i].scope, fake.evalArgsForCall[i].opHandle
}

func (fake *fakeEvalArrayInitializerer) EvalReturns(result1 []interface{}, result2 error) {
	fake.EvalStub = nil
	fake.evalReturns = struct {
		result1 []interface{}
		result2 error
	}{result1, result2}
}

func (fake *fakeEvalArrayInitializerer) EvalReturnsOnCall(i int, result1 []interface{}, result2 error) {
	fake.EvalStub = nil
	if fake.evalReturnsOnCall == nil {
		fake.evalReturnsOnCall = make(map[int]struct {
			result1 []interface{}
			result2 error
		})
	}
	fake.evalReturnsOnCall[i] = struct {
		result1 []interface{}
		result2 error
	}{result1, result2}
}

func (fake *fakeEvalArrayInitializerer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evalMutex.RLock()
	defer fake.evalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeEvalArrayInitializerer) recordInvocation(key string, args []interface{}) {
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
