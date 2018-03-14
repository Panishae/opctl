// Code generated by counterfeiter. DO NOT EDIT.
package data

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakeCoercer struct {
	CoerceToArrayStub        func(value *model.Value) (*model.Value, error)
	coerceToArrayMutex       sync.RWMutex
	coerceToArrayArgsForCall []struct {
		value *model.Value
	}
	coerceToArrayReturns struct {
		result1 *model.Value
		result2 error
	}
	coerceToArrayReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	CoerceToBooleanStub        func(value *model.Value) (*model.Value, error)
	coerceToBooleanMutex       sync.RWMutex
	coerceToBooleanArgsForCall []struct {
		value *model.Value
	}
	coerceToBooleanReturns struct {
		result1 *model.Value
		result2 error
	}
	coerceToBooleanReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	CoerceToFileStub        func(value *model.Value, scratchDir string) (*model.Value, error)
	coerceToFileMutex       sync.RWMutex
	coerceToFileArgsForCall []struct {
		value      *model.Value
		scratchDir string
	}
	coerceToFileReturns struct {
		result1 *model.Value
		result2 error
	}
	coerceToFileReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	CoerceToNumberStub        func(value *model.Value) (*model.Value, error)
	coerceToNumberMutex       sync.RWMutex
	coerceToNumberArgsForCall []struct {
		value *model.Value
	}
	coerceToNumberReturns struct {
		result1 *model.Value
		result2 error
	}
	coerceToNumberReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	CoerceToObjectStub        func(value *model.Value) (*model.Value, error)
	coerceToObjectMutex       sync.RWMutex
	coerceToObjectArgsForCall []struct {
		value *model.Value
	}
	coerceToObjectReturns struct {
		result1 *model.Value
		result2 error
	}
	coerceToObjectReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	CoerceToStringStub        func(value *model.Value) (*model.Value, error)
	coerceToStringMutex       sync.RWMutex
	coerceToStringArgsForCall []struct {
		value *model.Value
	}
	coerceToStringReturns struct {
		result1 *model.Value
		result2 error
	}
	coerceToStringReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeCoercer) CoerceToArray(value *model.Value) (*model.Value, error) {
	fake.coerceToArrayMutex.Lock()
	ret, specificReturn := fake.coerceToArrayReturnsOnCall[len(fake.coerceToArrayArgsForCall)]
	fake.coerceToArrayArgsForCall = append(fake.coerceToArrayArgsForCall, struct {
		value *model.Value
	}{value})
	fake.recordInvocation("CoerceToArray", []interface{}{value})
	fake.coerceToArrayMutex.Unlock()
	if fake.CoerceToArrayStub != nil {
		return fake.CoerceToArrayStub(value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.coerceToArrayReturns.result1, fake.coerceToArrayReturns.result2
}

func (fake *fakeCoercer) CoerceToArrayCallCount() int {
	fake.coerceToArrayMutex.RLock()
	defer fake.coerceToArrayMutex.RUnlock()
	return len(fake.coerceToArrayArgsForCall)
}

func (fake *fakeCoercer) CoerceToArrayArgsForCall(i int) *model.Value {
	fake.coerceToArrayMutex.RLock()
	defer fake.coerceToArrayMutex.RUnlock()
	return fake.coerceToArrayArgsForCall[i].value
}

func (fake *fakeCoercer) CoerceToArrayReturns(result1 *model.Value, result2 error) {
	fake.CoerceToArrayStub = nil
	fake.coerceToArrayReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToArrayReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.CoerceToArrayStub = nil
	if fake.coerceToArrayReturnsOnCall == nil {
		fake.coerceToArrayReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.coerceToArrayReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToBoolean(value *model.Value) (*model.Value, error) {
	fake.coerceToBooleanMutex.Lock()
	ret, specificReturn := fake.coerceToBooleanReturnsOnCall[len(fake.coerceToBooleanArgsForCall)]
	fake.coerceToBooleanArgsForCall = append(fake.coerceToBooleanArgsForCall, struct {
		value *model.Value
	}{value})
	fake.recordInvocation("CoerceToBoolean", []interface{}{value})
	fake.coerceToBooleanMutex.Unlock()
	if fake.CoerceToBooleanStub != nil {
		return fake.CoerceToBooleanStub(value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.coerceToBooleanReturns.result1, fake.coerceToBooleanReturns.result2
}

func (fake *fakeCoercer) CoerceToBooleanCallCount() int {
	fake.coerceToBooleanMutex.RLock()
	defer fake.coerceToBooleanMutex.RUnlock()
	return len(fake.coerceToBooleanArgsForCall)
}

func (fake *fakeCoercer) CoerceToBooleanArgsForCall(i int) *model.Value {
	fake.coerceToBooleanMutex.RLock()
	defer fake.coerceToBooleanMutex.RUnlock()
	return fake.coerceToBooleanArgsForCall[i].value
}

func (fake *fakeCoercer) CoerceToBooleanReturns(result1 *model.Value, result2 error) {
	fake.CoerceToBooleanStub = nil
	fake.coerceToBooleanReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToBooleanReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.CoerceToBooleanStub = nil
	if fake.coerceToBooleanReturnsOnCall == nil {
		fake.coerceToBooleanReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.coerceToBooleanReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToFile(value *model.Value, scratchDir string) (*model.Value, error) {
	fake.coerceToFileMutex.Lock()
	ret, specificReturn := fake.coerceToFileReturnsOnCall[len(fake.coerceToFileArgsForCall)]
	fake.coerceToFileArgsForCall = append(fake.coerceToFileArgsForCall, struct {
		value      *model.Value
		scratchDir string
	}{value, scratchDir})
	fake.recordInvocation("CoerceToFile", []interface{}{value, scratchDir})
	fake.coerceToFileMutex.Unlock()
	if fake.CoerceToFileStub != nil {
		return fake.CoerceToFileStub(value, scratchDir)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.coerceToFileReturns.result1, fake.coerceToFileReturns.result2
}

func (fake *fakeCoercer) CoerceToFileCallCount() int {
	fake.coerceToFileMutex.RLock()
	defer fake.coerceToFileMutex.RUnlock()
	return len(fake.coerceToFileArgsForCall)
}

func (fake *fakeCoercer) CoerceToFileArgsForCall(i int) (*model.Value, string) {
	fake.coerceToFileMutex.RLock()
	defer fake.coerceToFileMutex.RUnlock()
	return fake.coerceToFileArgsForCall[i].value, fake.coerceToFileArgsForCall[i].scratchDir
}

func (fake *fakeCoercer) CoerceToFileReturns(result1 *model.Value, result2 error) {
	fake.CoerceToFileStub = nil
	fake.coerceToFileReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToFileReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.CoerceToFileStub = nil
	if fake.coerceToFileReturnsOnCall == nil {
		fake.coerceToFileReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.coerceToFileReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToNumber(value *model.Value) (*model.Value, error) {
	fake.coerceToNumberMutex.Lock()
	ret, specificReturn := fake.coerceToNumberReturnsOnCall[len(fake.coerceToNumberArgsForCall)]
	fake.coerceToNumberArgsForCall = append(fake.coerceToNumberArgsForCall, struct {
		value *model.Value
	}{value})
	fake.recordInvocation("CoerceToNumber", []interface{}{value})
	fake.coerceToNumberMutex.Unlock()
	if fake.CoerceToNumberStub != nil {
		return fake.CoerceToNumberStub(value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.coerceToNumberReturns.result1, fake.coerceToNumberReturns.result2
}

func (fake *fakeCoercer) CoerceToNumberCallCount() int {
	fake.coerceToNumberMutex.RLock()
	defer fake.coerceToNumberMutex.RUnlock()
	return len(fake.coerceToNumberArgsForCall)
}

func (fake *fakeCoercer) CoerceToNumberArgsForCall(i int) *model.Value {
	fake.coerceToNumberMutex.RLock()
	defer fake.coerceToNumberMutex.RUnlock()
	return fake.coerceToNumberArgsForCall[i].value
}

func (fake *fakeCoercer) CoerceToNumberReturns(result1 *model.Value, result2 error) {
	fake.CoerceToNumberStub = nil
	fake.coerceToNumberReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToNumberReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.CoerceToNumberStub = nil
	if fake.coerceToNumberReturnsOnCall == nil {
		fake.coerceToNumberReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.coerceToNumberReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToObject(value *model.Value) (*model.Value, error) {
	fake.coerceToObjectMutex.Lock()
	ret, specificReturn := fake.coerceToObjectReturnsOnCall[len(fake.coerceToObjectArgsForCall)]
	fake.coerceToObjectArgsForCall = append(fake.coerceToObjectArgsForCall, struct {
		value *model.Value
	}{value})
	fake.recordInvocation("CoerceToObject", []interface{}{value})
	fake.coerceToObjectMutex.Unlock()
	if fake.CoerceToObjectStub != nil {
		return fake.CoerceToObjectStub(value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.coerceToObjectReturns.result1, fake.coerceToObjectReturns.result2
}

func (fake *fakeCoercer) CoerceToObjectCallCount() int {
	fake.coerceToObjectMutex.RLock()
	defer fake.coerceToObjectMutex.RUnlock()
	return len(fake.coerceToObjectArgsForCall)
}

func (fake *fakeCoercer) CoerceToObjectArgsForCall(i int) *model.Value {
	fake.coerceToObjectMutex.RLock()
	defer fake.coerceToObjectMutex.RUnlock()
	return fake.coerceToObjectArgsForCall[i].value
}

func (fake *fakeCoercer) CoerceToObjectReturns(result1 *model.Value, result2 error) {
	fake.CoerceToObjectStub = nil
	fake.coerceToObjectReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToObjectReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.CoerceToObjectStub = nil
	if fake.coerceToObjectReturnsOnCall == nil {
		fake.coerceToObjectReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.coerceToObjectReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToString(value *model.Value) (*model.Value, error) {
	fake.coerceToStringMutex.Lock()
	ret, specificReturn := fake.coerceToStringReturnsOnCall[len(fake.coerceToStringArgsForCall)]
	fake.coerceToStringArgsForCall = append(fake.coerceToStringArgsForCall, struct {
		value *model.Value
	}{value})
	fake.recordInvocation("CoerceToString", []interface{}{value})
	fake.coerceToStringMutex.Unlock()
	if fake.CoerceToStringStub != nil {
		return fake.CoerceToStringStub(value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.coerceToStringReturns.result1, fake.coerceToStringReturns.result2
}

func (fake *fakeCoercer) CoerceToStringCallCount() int {
	fake.coerceToStringMutex.RLock()
	defer fake.coerceToStringMutex.RUnlock()
	return len(fake.coerceToStringArgsForCall)
}

func (fake *fakeCoercer) CoerceToStringArgsForCall(i int) *model.Value {
	fake.coerceToStringMutex.RLock()
	defer fake.coerceToStringMutex.RUnlock()
	return fake.coerceToStringArgsForCall[i].value
}

func (fake *fakeCoercer) CoerceToStringReturns(result1 *model.Value, result2 error) {
	fake.CoerceToStringStub = nil
	fake.coerceToStringReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) CoerceToStringReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.CoerceToStringStub = nil
	if fake.coerceToStringReturnsOnCall == nil {
		fake.coerceToStringReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.coerceToStringReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *fakeCoercer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.coerceToArrayMutex.RLock()
	defer fake.coerceToArrayMutex.RUnlock()
	fake.coerceToBooleanMutex.RLock()
	defer fake.coerceToBooleanMutex.RUnlock()
	fake.coerceToFileMutex.RLock()
	defer fake.coerceToFileMutex.RUnlock()
	fake.coerceToNumberMutex.RLock()
	defer fake.coerceToNumberMutex.RUnlock()
	fake.coerceToObjectMutex.RLock()
	defer fake.coerceToObjectMutex.RUnlock()
	fake.coerceToStringMutex.RLock()
	defer fake.coerceToStringMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeCoercer) recordInvocation(key string, args []interface{}) {
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
