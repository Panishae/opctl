// Code generated by counterfeiter. DO NOT EDIT.
package data

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/api/client"
)

type Fake struct {
	NewFSProviderStub        func(...string) Provider
	newFSProviderMutex       sync.RWMutex
	newFSProviderArgsForCall []struct {
		arg1 []string
	}
	newFSProviderReturns struct {
		result1 Provider
	}
	newFSProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	NewGitProviderStub        func(string, *model.PullCreds) Provider
	newGitProviderMutex       sync.RWMutex
	newGitProviderArgsForCall []struct {
		arg1 string
		arg2 *model.PullCreds
	}
	newGitProviderReturns struct {
		result1 Provider
	}
	newGitProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	NewNodeProviderStub        func(client.Client, *model.PullCreds) Provider
	newNodeProviderMutex       sync.RWMutex
	newNodeProviderArgsForCall []struct {
		arg1 client.Client
		arg2 *model.PullCreds
	}
	newNodeProviderReturns struct {
		result1 Provider
	}
	newNodeProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	ResolveStub        func(context.Context, string, ...Provider) (model.DataHandle, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []Provider
	}
	resolveReturns struct {
		result1 model.DataHandle
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 model.DataHandle
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) NewFSProvider(arg1 ...string) Provider {
	fake.newFSProviderMutex.Lock()
	ret, specificReturn := fake.newFSProviderReturnsOnCall[len(fake.newFSProviderArgsForCall)]
	fake.newFSProviderArgsForCall = append(fake.newFSProviderArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("NewFSProvider", []interface{}{arg1})
	fake.newFSProviderMutex.Unlock()
	if fake.NewFSProviderStub != nil {
		return fake.NewFSProviderStub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newFSProviderReturns
	return fakeReturns.result1
}

func (fake *Fake) NewFSProviderCallCount() int {
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	return len(fake.newFSProviderArgsForCall)
}

func (fake *Fake) NewFSProviderCalls(stub func(...string) Provider) {
	fake.newFSProviderMutex.Lock()
	defer fake.newFSProviderMutex.Unlock()
	fake.NewFSProviderStub = stub
}

func (fake *Fake) NewFSProviderArgsForCall(i int) []string {
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	argsForCall := fake.newFSProviderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Fake) NewFSProviderReturns(result1 Provider) {
	fake.newFSProviderMutex.Lock()
	defer fake.newFSProviderMutex.Unlock()
	fake.NewFSProviderStub = nil
	fake.newFSProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewFSProviderReturnsOnCall(i int, result1 Provider) {
	fake.newFSProviderMutex.Lock()
	defer fake.newFSProviderMutex.Unlock()
	fake.NewFSProviderStub = nil
	if fake.newFSProviderReturnsOnCall == nil {
		fake.newFSProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newFSProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewGitProvider(arg1 string, arg2 *model.PullCreds) Provider {
	fake.newGitProviderMutex.Lock()
	ret, specificReturn := fake.newGitProviderReturnsOnCall[len(fake.newGitProviderArgsForCall)]
	fake.newGitProviderArgsForCall = append(fake.newGitProviderArgsForCall, struct {
		arg1 string
		arg2 *model.PullCreds
	}{arg1, arg2})
	fake.recordInvocation("NewGitProvider", []interface{}{arg1, arg2})
	fake.newGitProviderMutex.Unlock()
	if fake.NewGitProviderStub != nil {
		return fake.NewGitProviderStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newGitProviderReturns
	return fakeReturns.result1
}

func (fake *Fake) NewGitProviderCallCount() int {
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	return len(fake.newGitProviderArgsForCall)
}

func (fake *Fake) NewGitProviderCalls(stub func(string, *model.PullCreds) Provider) {
	fake.newGitProviderMutex.Lock()
	defer fake.newGitProviderMutex.Unlock()
	fake.NewGitProviderStub = stub
}

func (fake *Fake) NewGitProviderArgsForCall(i int) (string, *model.PullCreds) {
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	argsForCall := fake.newGitProviderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Fake) NewGitProviderReturns(result1 Provider) {
	fake.newGitProviderMutex.Lock()
	defer fake.newGitProviderMutex.Unlock()
	fake.NewGitProviderStub = nil
	fake.newGitProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewGitProviderReturnsOnCall(i int, result1 Provider) {
	fake.newGitProviderMutex.Lock()
	defer fake.newGitProviderMutex.Unlock()
	fake.NewGitProviderStub = nil
	if fake.newGitProviderReturnsOnCall == nil {
		fake.newGitProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newGitProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewNodeProvider(arg1 client.Client, arg2 *model.PullCreds) Provider {
	fake.newNodeProviderMutex.Lock()
	ret, specificReturn := fake.newNodeProviderReturnsOnCall[len(fake.newNodeProviderArgsForCall)]
	fake.newNodeProviderArgsForCall = append(fake.newNodeProviderArgsForCall, struct {
		arg1 client.Client
		arg2 *model.PullCreds
	}{arg1, arg2})
	fake.recordInvocation("NewNodeProvider", []interface{}{arg1, arg2})
	fake.newNodeProviderMutex.Unlock()
	if fake.NewNodeProviderStub != nil {
		return fake.NewNodeProviderStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newNodeProviderReturns
	return fakeReturns.result1
}

func (fake *Fake) NewNodeProviderCallCount() int {
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	return len(fake.newNodeProviderArgsForCall)
}

func (fake *Fake) NewNodeProviderCalls(stub func(client.Client, *model.PullCreds) Provider) {
	fake.newNodeProviderMutex.Lock()
	defer fake.newNodeProviderMutex.Unlock()
	fake.NewNodeProviderStub = stub
}

func (fake *Fake) NewNodeProviderArgsForCall(i int) (client.Client, *model.PullCreds) {
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	argsForCall := fake.newNodeProviderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Fake) NewNodeProviderReturns(result1 Provider) {
	fake.newNodeProviderMutex.Lock()
	defer fake.newNodeProviderMutex.Unlock()
	fake.NewNodeProviderStub = nil
	fake.newNodeProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewNodeProviderReturnsOnCall(i int, result1 Provider) {
	fake.newNodeProviderMutex.Lock()
	defer fake.newNodeProviderMutex.Unlock()
	fake.NewNodeProviderStub = nil
	if fake.newNodeProviderReturnsOnCall == nil {
		fake.newNodeProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newNodeProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) Resolve(arg1 context.Context, arg2 string, arg3 ...Provider) (model.DataHandle, error) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []Provider
	}{arg1, arg2, arg3})
	fake.recordInvocation("Resolve", []interface{}{arg1, arg2, arg3})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Fake) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *Fake) ResolveCalls(stub func(context.Context, string, ...Provider) (model.DataHandle, error)) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = stub
}

func (fake *Fake) ResolveArgsForCall(i int) (context.Context, string, []Provider) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	argsForCall := fake.resolveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Fake) ResolveReturns(result1 model.DataHandle, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) ResolveReturnsOnCall(i int, result1 model.DataHandle, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 model.DataHandle
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
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

var _ Data = new(Fake)
