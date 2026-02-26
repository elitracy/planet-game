package orders

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/actions"
)

type Order struct {
	task.Task
	ID              task.TaskID
	Name            string
	Actions         []*actions.Action
	StartTick       engine.Tick
	Status          task.Status
	currentActionID task.TaskID
}

func (o Order) GetID() task.TaskID                { return o.ID }
func (o *Order) SetID(id task.TaskID)             { o.ID = id }
func (o Order) GetName() string                   { return o.Name }
func (o Order) GetActions() []*actions.Action     { return o.Actions }
func (o *Order) AddAction(action *actions.Action) { o.Actions = append(o.Actions, action) }
func (o Order) GetStartTick() engine.Tick         { return o.StartTick }
func (o Order) GetStatus() task.Status            { return o.Status }
func (o *Order) SetStatus(status task.Status)     { o.Status = status }

func (o Order) GetDuration() engine.Tick {
	var duration engine.Tick
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}
