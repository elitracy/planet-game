package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/actions"
)

type Order struct {
	events.Event
	ID              events.EventID
	Name            string
	Actions         []*actions.Action
	StartTick       core.Tick
	Status          events.EventStatus
	currentActionID events.EventID
}

func (o Order) GetID() events.EventID                { return o.ID }
func (o *Order) SetID(id events.EventID)             { o.ID = id }
func (o Order) GetName() string                      { return o.Name }
func (o Order) GetActions() []*actions.Action        { return o.Actions }
func (o *Order) AddAction(action *actions.Action)    { o.Actions = append(o.Actions, action) }
func (o Order) GetStartTick() core.Tick              { return o.StartTick }
func (o Order) GetStatus() events.EventStatus        { return o.Status }
func (o *Order) SetStatus(status events.EventStatus) { o.Status = status }

func (o Order) GetDuration() core.Tick {
	var duration core.Tick
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}
