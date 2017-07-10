// Code generated by counterfeiter. DO NOT EDIT.
package manifest

import (
	"io"
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	ValidateStub        func(manifestReader io.Reader) []error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		manifestReader io.Reader
	}
	validateReturns struct {
		result1 []error
	}
	validateReturnsOnCall map[int]struct {
		result1 []error
	}
	UnmarshalStub        func(manifestReader io.Reader) (*model.PkgManifest, error)
	unmarshalMutex       sync.RWMutex
	unmarshalArgsForCall []struct {
		manifestReader io.Reader
	}
	unmarshalReturns struct {
		result1 *model.PkgManifest
		result2 error
	}
	unmarshalReturnsOnCall map[int]struct {
		result1 *model.PkgManifest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Validate(manifestReader io.Reader) []error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		manifestReader io.Reader
	}{manifestReader})
	fake.recordInvocation("Validate", []interface{}{manifestReader})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(manifestReader)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateReturns.result1
}

func (fake *Fake) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *Fake) ValidateArgsForCall(i int) io.Reader {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].manifestReader
}

func (fake *Fake) ValidateReturns(result1 []error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) ValidateReturnsOnCall(i int, result1 []error) {
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

func (fake *Fake) Unmarshal(manifestReader io.Reader) (*model.PkgManifest, error) {
	fake.unmarshalMutex.Lock()
	ret, specificReturn := fake.unmarshalReturnsOnCall[len(fake.unmarshalArgsForCall)]
	fake.unmarshalArgsForCall = append(fake.unmarshalArgsForCall, struct {
		manifestReader io.Reader
	}{manifestReader})
	fake.recordInvocation("Unmarshal", []interface{}{manifestReader})
	fake.unmarshalMutex.Unlock()
	if fake.UnmarshalStub != nil {
		return fake.UnmarshalStub(manifestReader)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unmarshalReturns.result1, fake.unmarshalReturns.result2
}

func (fake *Fake) UnmarshalCallCount() int {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return len(fake.unmarshalArgsForCall)
}

func (fake *Fake) UnmarshalArgsForCall(i int) io.Reader {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return fake.unmarshalArgsForCall[i].manifestReader
}

func (fake *Fake) UnmarshalReturns(result1 *model.PkgManifest, result2 error) {
	fake.UnmarshalStub = nil
	fake.unmarshalReturns = struct {
		result1 *model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) UnmarshalReturnsOnCall(i int, result1 *model.PkgManifest, result2 error) {
	fake.UnmarshalStub = nil
	if fake.unmarshalReturnsOnCall == nil {
		fake.unmarshalReturnsOnCall = make(map[int]struct {
			result1 *model.PkgManifest
			result2 error
		})
	}
	fake.unmarshalReturnsOnCall[i] = struct {
		result1 *model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
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

var _ Manifest = new(Fake)
