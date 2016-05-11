package models

import (
  "github.com/dev-op-spec/engine/core/models"
  "reflect"
  "strings"
)

func NewEventMsg(
event models.Event,
) EventMsg {

  eventType := strings.TrimSuffix(
    reflect.TypeOf(event).Elem().Name(),
    "Event",
  )

  return EventMsg{
    Type:eventType,
    Data: event,
  }

}

/*
a msg for passing around an event
 */
type EventMsg struct {
  Type string `json:"type"`
  Data interface{} `json:"data"`
}