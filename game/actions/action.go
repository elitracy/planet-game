package actions

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/models"
)

type Action struct {
	engine.Event
	ID           engine.EventID
	TargetEntity models.Entity
	Description  string
	StartTick    engine.Tick
	Duration     engine.Tick
	Status       engine.EventStatus
}

func (a Action) GetID() engine.EventID                { return a.ID }
func (a *Action) SetID(id engine.EventID)             { a.ID = id }
func (a Action) GetTargetEntity() models.Entity       { return a.TargetEntity }
func (a Action) GetDescription() string               { return a.Description }
func (a Action) GetStartTick() engine.Tick            { return a.StartTick }
func (a Action) GetDuration() engine.Tick             { return a.Duration }
func (a Action) GetStatus() engine.EventStatus        { return a.Status }
func (a *Action) SetStatus(status engine.EventStatus) { a.Status = status }
func (a *Action) Execute()                            {}
