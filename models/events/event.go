package events

import (
	"github.com/elitracy/planets/core"
)

//go:generate stringer -type=EventStatus
type EventStatus int

const (
	EventPending EventStatus = iota
	EventExecuting
	EventComplete
	EventFailed
)

type Event interface {
	GetID() EventID
	SetID(EventID)
	GetExecuteTick() core.Tick
	GetDuration() core.Tick
	GetStatus() EventStatus
}
