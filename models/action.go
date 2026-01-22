package models

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
)

type Action interface {
	GetID() int
	GetTargetEntity() Entity
	GetDescription() string
	GetExecuteTick() core.Tick
	GetDuration() core.Tick
	GetStatus() consts.EventStatus
	SetStatus(consts.EventStatus)
	Execute()
}
