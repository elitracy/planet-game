package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/actions"
)

type Order struct {
	events.Event
	ID          int
	Name        string
	Actions     []actions.OrderAction
	ExecuteTick core.Tick
	Status      consts.EventStatus
}

func (o Order) GetID() int                        { return o.ID }
func (o *Order) SetID(id int)                     { o.ID = id }
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
