package actions

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/models"
)

type Action struct {
	task.Task
	ID           task.TaskID
	TargetEntity models.Entity
	Description  string
	StartTick    engine.Tick
	Duration     engine.Tick
	Status       task.Status
	Execute      func()
}

func (a Action) GetID() task.TaskID             { return a.ID }
func (a *Action) SetID(id task.TaskID)          { a.ID = id }
func (a Action) GetTargetEntity() models.Entity { return a.TargetEntity }
func (a Action) GetDescription() string         { return a.Description }
func (a Action) GetStartTick() engine.Tick      { return a.StartTick }
func (a Action) GetDuration() engine.Tick       { return a.Duration }
func (a Action) GetStatus() task.Status         { return a.Status }
func (a *Action) SetStatus(status task.Status)  { a.Status = status }
