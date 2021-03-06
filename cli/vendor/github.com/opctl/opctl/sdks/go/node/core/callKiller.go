package core

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./fakeCallKiller.go --fake-name fakeCallKiller ./ callKiller

import (
	"sync"

	"github.com/opctl/opctl/sdks/go/model"

	"github.com/opctl/opctl/sdks/go/node/core/containerruntime"
	"github.com/opctl/opctl/sdks/go/pubsub"
)

type callKiller interface {
	Kill(
		callID string,
		rootCallID string,
	)
}

func newCallKiller(
	callStore callStore,
	containerRuntime containerruntime.ContainerRuntime,
	eventPublisher pubsub.EventPublisher,
) callKiller {
	return _callKiller{
		callStore:        callStore,
		containerRuntime: containerRuntime,
		eventPublisher:   eventPublisher,
	}
}

type _callKiller struct {
	callStore        callStore
	containerRuntime containerruntime.ContainerRuntime
	eventPublisher   pubsub.EventPublisher
}

func (ckr _callKiller) Kill(
	callID string,
	rootCallID string,
) {
	ckr.eventPublisher.Publish(model.Event{
		CallKilled: &model.CallKilledEvent{
			CallID:     callID,
			RootCallID: rootCallID,
		},
	})

	ckr.callStore.SetIsKilled(callID)
	ckr.containerRuntime.DeleteContainerIfExists(callID)

	var waitGroup sync.WaitGroup

	for _, childCallGraph := range ckr.callStore.ListWithParentID(callID) {
		// recursively kill all child calls
		waitGroup.Add(1)
		go func(childCallGraph *model.DCG) {
			defer waitGroup.Done()

			ckr.Kill(
				childCallGraph.Id,
				rootCallID,
			)

		}(childCallGraph)
	}

	waitGroup.Wait()

}
