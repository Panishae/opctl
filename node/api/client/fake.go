// Code generated by counterfeiter. DO NOT EDIT.
package client

import (
	"context"
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	GetEventStreamStub        func(req *model.GetEventStreamReq) (stream chan model.Event, err error)
	getEventStreamMutex       sync.RWMutex
	getEventStreamArgsForCall []struct {
		req *model.GetEventStreamReq
	}
	getEventStreamReturns struct {
		result1 chan model.Event
		result2 error
	}
	getEventStreamReturnsOnCall map[int]struct {
		result1 chan model.Event
		result2 error
	}
	GetDataStub        func(ctx context.Context, req model.GetDataReq) (model.ReadSeekCloser, error)
	getDataMutex       sync.RWMutex
	getDataArgsForCall []struct {
		ctx context.Context
		req model.GetDataReq
	}
	getDataReturns struct {
		result1 model.ReadSeekCloser
		result2 error
	}
	getDataReturnsOnCall map[int]struct {
		result1 model.ReadSeekCloser
		result2 error
	}
	KillOpStub        func(ctx context.Context, req model.KillOpReq) (err error)
	killOpMutex       sync.RWMutex
	killOpArgsForCall []struct {
		ctx context.Context
		req model.KillOpReq
	}
	killOpReturns struct {
		result1 error
	}
	killOpReturnsOnCall map[int]struct {
		result1 error
	}
	ListDescendantsStub      func(ctx context.Context, req model.ListDescendantsReq) ([]*model.DirEntry, error)
	listDirEntrysMutex       sync.RWMutex
	listDirEntrysArgsForCall []struct {
		ctx context.Context
		req model.ListDescendantsReq
	}
	listDirEntrysReturns struct {
		result1 []*model.DirEntry
		result2 error
	}
	listDirEntrysReturnsOnCall map[int]struct {
		result1 []*model.DirEntry
		result2 error
	}
	StartOpStub        func(ctx context.Context, req model.StartOpReq) (opID string, err error)
	startOpMutex       sync.RWMutex
	startOpArgsForCall []struct {
		ctx context.Context
		req model.StartOpReq
	}
	startOpReturns struct {
		result1 string
		result2 error
	}
	startOpReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) GetEventStream(req *model.GetEventStreamReq) (stream chan model.Event, err error) {
	fake.getEventStreamMutex.Lock()
	ret, specificReturn := fake.getEventStreamReturnsOnCall[len(fake.getEventStreamArgsForCall)]
	fake.getEventStreamArgsForCall = append(fake.getEventStreamArgsForCall, struct {
		req *model.GetEventStreamReq
	}{req})
	fake.recordInvocation("GetEventStream", []interface{}{req})
	fake.getEventStreamMutex.Unlock()
	if fake.GetEventStreamStub != nil {
		return fake.GetEventStreamStub(req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getEventStreamReturns.result1, fake.getEventStreamReturns.result2
}

func (fake *Fake) GetEventStreamCallCount() int {
	fake.getEventStreamMutex.RLock()
	defer fake.getEventStreamMutex.RUnlock()
	return len(fake.getEventStreamArgsForCall)
}

func (fake *Fake) GetEventStreamArgsForCall(i int) *model.GetEventStreamReq {
	fake.getEventStreamMutex.RLock()
	defer fake.getEventStreamMutex.RUnlock()
	return fake.getEventStreamArgsForCall[i].req
}

func (fake *Fake) GetEventStreamReturns(result1 chan model.Event, result2 error) {
	fake.GetEventStreamStub = nil
	fake.getEventStreamReturns = struct {
		result1 chan model.Event
		result2 error
	}{result1, result2}
}

func (fake *Fake) GetEventStreamReturnsOnCall(i int, result1 chan model.Event, result2 error) {
	fake.GetEventStreamStub = nil
	if fake.getEventStreamReturnsOnCall == nil {
		fake.getEventStreamReturnsOnCall = make(map[int]struct {
			result1 chan model.Event
			result2 error
		})
	}
	fake.getEventStreamReturnsOnCall[i] = struct {
		result1 chan model.Event
		result2 error
	}{result1, result2}
}

func (fake *Fake) GetData(ctx context.Context, req model.GetDataReq) (model.ReadSeekCloser, error) {
	fake.getDataMutex.Lock()
	ret, specificReturn := fake.getDataReturnsOnCall[len(fake.getDataArgsForCall)]
	fake.getDataArgsForCall = append(fake.getDataArgsForCall, struct {
		ctx context.Context
		req model.GetDataReq
	}{ctx, req})
	fake.recordInvocation("GetData", []interface{}{ctx, req})
	fake.getDataMutex.Unlock()
	if fake.GetDataStub != nil {
		return fake.GetDataStub(ctx, req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDataReturns.result1, fake.getDataReturns.result2
}

func (fake *Fake) GetDataCallCount() int {
	fake.getDataMutex.RLock()
	defer fake.getDataMutex.RUnlock()
	return len(fake.getDataArgsForCall)
}

func (fake *Fake) GetDataArgsForCall(i int) (context.Context, model.GetDataReq) {
	fake.getDataMutex.RLock()
	defer fake.getDataMutex.RUnlock()
	return fake.getDataArgsForCall[i].ctx, fake.getDataArgsForCall[i].req
}

func (fake *Fake) GetDataReturns(result1 model.ReadSeekCloser, result2 error) {
	fake.GetDataStub = nil
	fake.getDataReturns = struct {
		result1 model.ReadSeekCloser
		result2 error
	}{result1, result2}
}

func (fake *Fake) GetDataReturnsOnCall(i int, result1 model.ReadSeekCloser, result2 error) {
	fake.GetDataStub = nil
	if fake.getDataReturnsOnCall == nil {
		fake.getDataReturnsOnCall = make(map[int]struct {
			result1 model.ReadSeekCloser
			result2 error
		})
	}
	fake.getDataReturnsOnCall[i] = struct {
		result1 model.ReadSeekCloser
		result2 error
	}{result1, result2}
}

func (fake *Fake) KillOp(ctx context.Context, req model.KillOpReq) (err error) {
	fake.killOpMutex.Lock()
	ret, specificReturn := fake.killOpReturnsOnCall[len(fake.killOpArgsForCall)]
	fake.killOpArgsForCall = append(fake.killOpArgsForCall, struct {
		ctx context.Context
		req model.KillOpReq
	}{ctx, req})
	fake.recordInvocation("KillOp", []interface{}{ctx, req})
	fake.killOpMutex.Unlock()
	if fake.KillOpStub != nil {
		return fake.KillOpStub(ctx, req)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.killOpReturns.result1
}

func (fake *Fake) KillOpCallCount() int {
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	return len(fake.killOpArgsForCall)
}

func (fake *Fake) KillOpArgsForCall(i int) (context.Context, model.KillOpReq) {
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	return fake.killOpArgsForCall[i].ctx, fake.killOpArgsForCall[i].req
}

func (fake *Fake) KillOpReturns(result1 error) {
	fake.KillOpStub = nil
	fake.killOpReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) KillOpReturnsOnCall(i int, result1 error) {
	fake.KillOpStub = nil
	if fake.killOpReturnsOnCall == nil {
		fake.killOpReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.killOpReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) ListDescendants(ctx context.Context, req model.ListDescendantsReq) ([]*model.DirEntry, error) {
	fake.listDirEntrysMutex.Lock()
	ret, specificReturn := fake.listDirEntrysReturnsOnCall[len(fake.listDirEntrysArgsForCall)]
	fake.listDirEntrysArgsForCall = append(fake.listDirEntrysArgsForCall, struct {
		ctx context.Context
		req model.ListDescendantsReq
	}{ctx, req})
	fake.recordInvocation("ListDescendants", []interface{}{ctx, req})
	fake.listDirEntrysMutex.Unlock()
	if fake.ListDescendantsStub != nil {
		return fake.ListDescendantsStub(ctx, req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDirEntrysReturns.result1, fake.listDirEntrysReturns.result2
}

func (fake *Fake) ListDescendantsCallCount() int {
	fake.listDirEntrysMutex.RLock()
	defer fake.listDirEntrysMutex.RUnlock()
	return len(fake.listDirEntrysArgsForCall)
}

func (fake *Fake) ListDescendantsArgsForCall(i int) (context.Context, model.ListDescendantsReq) {
	fake.listDirEntrysMutex.RLock()
	defer fake.listDirEntrysMutex.RUnlock()
	return fake.listDirEntrysArgsForCall[i].ctx, fake.listDirEntrysArgsForCall[i].req
}

func (fake *Fake) ListDescendantsReturns(result1 []*model.DirEntry, result2 error) {
	fake.ListDescendantsStub = nil
	fake.listDirEntrysReturns = struct {
		result1 []*model.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *Fake) ListDescendantsReturnsOnCall(i int, result1 []*model.DirEntry, result2 error) {
	fake.ListDescendantsStub = nil
	if fake.listDirEntrysReturnsOnCall == nil {
		fake.listDirEntrysReturnsOnCall = make(map[int]struct {
			result1 []*model.DirEntry
			result2 error
		})
	}
	fake.listDirEntrysReturnsOnCall[i] = struct {
		result1 []*model.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *Fake) StartOp(ctx context.Context, req model.StartOpReq) (opID string, err error) {
	fake.startOpMutex.Lock()
	ret, specificReturn := fake.startOpReturnsOnCall[len(fake.startOpArgsForCall)]
	fake.startOpArgsForCall = append(fake.startOpArgsForCall, struct {
		ctx context.Context
		req model.StartOpReq
	}{ctx, req})
	fake.recordInvocation("StartOp", []interface{}{ctx, req})
	fake.startOpMutex.Unlock()
	if fake.StartOpStub != nil {
		return fake.StartOpStub(ctx, req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startOpReturns.result1, fake.startOpReturns.result2
}

func (fake *Fake) StartOpCallCount() int {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	return len(fake.startOpArgsForCall)
}

func (fake *Fake) StartOpArgsForCall(i int) (context.Context, model.StartOpReq) {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	return fake.startOpArgsForCall[i].ctx, fake.startOpArgsForCall[i].req
}

func (fake *Fake) StartOpReturns(result1 string, result2 error) {
	fake.StartOpStub = nil
	fake.startOpReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) StartOpReturnsOnCall(i int, result1 string, result2 error) {
	fake.StartOpStub = nil
	if fake.startOpReturnsOnCall == nil {
		fake.startOpReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.startOpReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEventStreamMutex.RLock()
	defer fake.getEventStreamMutex.RUnlock()
	fake.getDataMutex.RLock()
	defer fake.getDataMutex.RUnlock()
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	fake.listDirEntrysMutex.RLock()
	defer fake.listDirEntrysMutex.RUnlock()
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
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

var _ Client = new(Fake)
