package core

//go:generate counterfeiter -o ./fakeCompositionRoot.go --fake-name fakeCompositionRoot ./ compositionRoot

import (
  "github.com/opctl/engine/core/ports"
  "github.com/opctl/engine/core/models"
)

type compositionRoot interface {
  RunOpUseCase() runOpUseCase
  AddSubOpUseCase() addSubOpUseCase
  GetEventStreamUseCase() getEventStreamUseCase
  KillOpRunUseCase() killOpRunUseCase
  ListOpsUseCase() listOpsUseCase
}

func newCompositionRoot(
containerEngine ports.ContainerEngine,
filesys ports.Filesys,
) (compositionRoot compositionRoot, err error) {

  /* factories */
  pathToOpsDirFactory := newPathToOpsDirFactory()

  pathToOpDirFactory := newPathToOpDirFactory(pathToOpsDirFactory)

  pathToOpFileFactory := newPathToOpFileFactory(pathToOpDirFactory)

  uniqueStringFactory := newUniqueStringFactory()

  /* components */
  eventStream := newEventStream()

  yamlCodec := newYamlCodec()

  logger := func(logEntryEmittedEvent models.LogEntryEmittedEvent) {
    eventStream.Publish(logEntryEmittedEvent)
  }

  opRunner := newOpRunner(
    containerEngine,
    eventStream,
    filesys,
    logger,
    uniqueStringFactory,
    yamlCodec,
  )

  /* use cases */
  runOpUseCase := newRunOpUseCase(
    opRunner,
    uniqueStringFactory,
  )

  addSubOpUseCase := newAddSubOpUseCase(
    filesys,
    pathToOpFileFactory,
    yamlCodec,
  )

  getEventStreamUseCase := newGetEventStreamUseCase(
    eventStream,
  )

  killOpRunUseCase := newKillOpRunUseCase(
    opRunner,
    uniqueStringFactory,
  )

  listOpsUseCase := newListOpsUseCase(
    filesys,
    pathToOpFileFactory,
    pathToOpsDirFactory,
    yamlCodec,
  )

  compositionRoot = &_compositionRoot{
    runOpUseCase: runOpUseCase,
    addSubOpUseCase: addSubOpUseCase,
    getEventStreamUseCase:getEventStreamUseCase,
    killOpRunUseCase:killOpRunUseCase,
    listOpsUseCase: listOpsUseCase,
  }

  return

}

type _compositionRoot struct {
  runOpUseCase              runOpUseCase
  addSubOpUseCase           addSubOpUseCase
  getEventStreamUseCase     getEventStreamUseCase
  killOpRunUseCase          killOpRunUseCase
  listOpsUseCase            listOpsUseCase
}

func (this _compositionRoot) RunOpUseCase() runOpUseCase {
  return this.runOpUseCase
}

func (this _compositionRoot) AddSubOpUseCase() addSubOpUseCase {
  return this.addSubOpUseCase
}

func (this _compositionRoot) GetEventStreamUseCase() getEventStreamUseCase {
  return this.getEventStreamUseCase
}

func (this _compositionRoot) KillOpRunUseCase() killOpRunUseCase {
  return this.killOpRunUseCase
}

func (this _compositionRoot) ListOpsUseCase() listOpsUseCase {
  return this.listOpsUseCase
}
