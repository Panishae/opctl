// Code generated by counterfeiter. DO NOT EDIT.
package pkg

import (
	"sync"
)

type fakeRefParser struct {
	ParseStub        func(pkgRef string) (*Ref, error)
	parseMutex       sync.RWMutex
	parseArgsForCall []struct {
		pkgRef string
	}
	parseReturns struct {
		result1 *Ref
		result2 error
	}
	parseReturnsOnCall map[int]struct {
		result1 *Ref
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeRefParser) Parse(pkgRef string) (*Ref, error) {
	fake.parseMutex.Lock()
	ret, specificReturn := fake.parseReturnsOnCall[len(fake.parseArgsForCall)]
	fake.parseArgsForCall = append(fake.parseArgsForCall, struct {
		pkgRef string
	}{pkgRef})
	fake.recordInvocation("Parse", []interface{}{pkgRef})
	fake.parseMutex.Unlock()
	if fake.ParseStub != nil {
		return fake.ParseStub(pkgRef)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.parseReturns.result1, fake.parseReturns.result2
}

func (fake *fakeRefParser) ParseCallCount() int {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return len(fake.parseArgsForCall)
}

func (fake *fakeRefParser) ParseArgsForCall(i int) string {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.parseArgsForCall[i].pkgRef
}

func (fake *fakeRefParser) ParseReturns(result1 *Ref, result2 error) {
	fake.ParseStub = nil
	fake.parseReturns = struct {
		result1 *Ref
		result2 error
	}{result1, result2}
}

func (fake *fakeRefParser) ParseReturnsOnCall(i int, result1 *Ref, result2 error) {
	fake.ParseStub = nil
	if fake.parseReturnsOnCall == nil {
		fake.parseReturnsOnCall = make(map[int]struct {
			result1 *Ref
			result2 error
		})
	}
	fake.parseReturnsOnCall[i] = struct {
		result1 *Ref
		result2 error
	}{result1, result2}
}

func (fake *fakeRefParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeRefParser) recordInvocation(key string, args []interface{}) {
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
