package interfaces

import (
	"github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/consts"
)

type Event interface {
	GetID() int
	GetExecuteTick() core.Tick
	GetDuration() core.Tick
	GetStatus() EventStatus
}
