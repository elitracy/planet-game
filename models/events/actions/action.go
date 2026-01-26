package actions

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
)

type Action struct {
	events.Event
	ID           events.EventID
	TargetEntity models.Entity
	Description  string
	ExecuteTick  core.Tick
	Duration     core.Tick
	Status       consts.EventStatus
}

func (a Action) GetID() events.EventID                { return a.ID }
func (a *Action) SetID(id events.EventID)             { a.ID = id }
func (a Action) GetTargetEntity() models.Entity       { return a.TargetEntity }
func (a Action) GetDescription() string               { return a.Description }
func (a Action) GetExecuteTick() core.Tick            { return a.ExecuteTick }
func (a Action) GetDuration() core.Tick               { return a.Duration }
func (a Action) GetStatus() consts.EventStatus        { return a.Status }
func (a *Action) SetStatus(status consts.EventStatus) { a.Status = status }

type OrderAction interface {
	GetID() events.EventID
	SetID(events.EventID)
	GetTargetEntity() models.Entity
	GetDescription() string
	GetExecuteTick() core.Tick
	GetDuration() core.Tick
	GetStatus() consts.EventStatus
	SetStatus(status consts.EventStatus)
	Execute()
}
