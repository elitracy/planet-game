package events

import (
	"github.com/elitracy/planets/core"
)

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

func (e EventStatus) String() string {
	switch e {
	case EventPending:
		return "Pending"
	case EventExecuting:
		return "Executing"
	case EventComplete:
		return "Complete"
	case EventFailed:
		return "Failed"
	default:
		return ""
	}
}
