package orders

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/actions"
)

type Order struct {
	engine.Event
	ID              engine.EventID
	Name            string
	Actions         []*actions.Action
	StartTick       engine.Tick
	Status          engine.EventStatus
	currentActionID engine.EventID
}

func (o Order) GetID() engine.EventID                { return o.ID }
func (o *Order) SetID(id engine.EventID)             { o.ID = id }
func (o Order) GetName() string                      { return o.Name }
func (o Order) GetActions() []*actions.Action        { return o.Actions }
func (o *Order) AddAction(action *actions.Action)    { o.Actions = append(o.Actions, action) }
func (o Order) GetStartTick() engine.Tick            { return o.StartTick }
func (o Order) GetStatus() engine.EventStatus        { return o.Status }
func (o *Order) SetStatus(status engine.EventStatus) { o.Status = status }

func (o Order) GetDuration() engine.Tick {
	var duration engine.Tick
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}
