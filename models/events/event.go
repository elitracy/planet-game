package events

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
)

type Event interface {
	GetID() EventID
	SetID(EventID)
	GetExecuteTick() core.Tick
	GetDuration() core.Tick
	GetStatus() consts.EventStatus
}
