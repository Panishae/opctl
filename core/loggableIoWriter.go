package core

import (
  "github.com/dev-op-spec/engine/core/models"
  "time"
  "strings"
  "io"
)

type LoggableIoWriter io.Writer

func NewLoggableIoWriter(
logChannel chan *models.LogEntry,
logEntryStream string,
) LoggableIoWriter {

  return &logEmittingIoWriter{
    logChannel:logChannel,
    logEntryStream:logEntryStream,
  }

}

type logEmittingIoWriter struct {
  logChannel     chan *models.LogEntry
  logEntryStream string
}

func (this logEmittingIoWriter) Write(
p []byte,
) (n int, err error) {

  n = len(p)

  this.logChannel <- models.NewLogEntry(
    strings.TrimSpace(string(p)),
    time.Now(),
    this.logEntryStream,
  )

  return

}