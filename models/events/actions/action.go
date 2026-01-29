package actions

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
)

type Action struct {
	events.Event
	ID           events.EventID
	TargetEntity models.Entity
	Description  string
	StartTick    core.Tick
	Duration     core.Tick
	Status       events.EventStatus
}

func (a Action) GetID() events.EventID                { return a.ID }
func (a *Action) SetID(id events.EventID)             { a.ID = id }
func (a Action) GetTargetEntity() models.Entity       { return a.TargetEntity }
func (a Action) GetDescription() string               { return a.Description }
func (a Action) GetStartTick() core.Tick              { return a.StartTick }
func (a Action) GetDuration() core.Tick               { return a.Duration }
func (a Action) GetStatus() events.EventStatus        { return a.Status }
func (a *Action) SetStatus(status events.EventStatus) { a.Status = status }
func (a *Action) Execute()                            {}
