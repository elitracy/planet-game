package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/actions"
)

type Order struct {
	events.Event
	ID          events.EventID
	Name        string
	Actions     []actions.OrderAction
	ExecuteTick core.Tick
	Status      consts.EventStatus
}

func (o Order) GetID() events.EventID             { return o.ID }
func (o *Order) SetID(id events.EventID)          { o.ID = id }
func (o Order) GetName() string                   { return o.Name }
func (o Order) GetActions() []actions.OrderAction { return o.Actions }
func (o Order) GetExecuteTick() core.Tick         { return o.ExecuteTick }
func (o Order) GetStatus() consts.EventStatus     { return o.Status }

func (o Order) GetDuration() core.Tick {
	var duration core.Tick
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}

func (o *Order) SetStatus(status consts.EventStatus) { o.Status = status }

func (o Order) GetEndTick() core.Tick {
	var latestTick core.Tick
	for _, action := range o.Actions {
		endTick := action.GetExecuteTick() + action.GetDuration()
		if endTick >= latestTick {
			latestTick = endTick
		}
	}

	return latestTick

}
