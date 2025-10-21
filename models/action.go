package models

import (
	"github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/consts"
	. "github.com/elitracy/planets/core/interfaces"
)

type Action interface {
	GetID() int
	GetTargetEntity() Entity
	GetDescription() string
	GetExecuteTick() core.Tick
	GetDuration() core.Tick
	GetStatus() EventStatus
	SetStatus(EventStatus)
	Execute()
}
