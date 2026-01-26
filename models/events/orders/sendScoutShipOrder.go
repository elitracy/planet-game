package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/actions"
)

type ScoutPositionOrder struct {
	*Order
	ship        *models.Ship
	destination models.Destination
}

func NewScoutDestinationOrder(ship *models.Ship, dest models.Destination, execTick core.Tick) *Order {
	name := ""
	if dest.Entity != nil {
		name = dest.Entity.GetName()
	} else {
		name = dest.Entity.GetPosition().String()
	}
	order := &ScoutPositionOrder{
		Order: &Order{
			Name:        "Scout " + name,
			ExecuteTick: execTick,
			Status:      events.EventPending,
		},
		ship:        ship,
		destination: dest,
	}

	d := core.EuclidianDistance(ship.Position, dest.Position)
	t := d / ship.Velocity.Vector()

	initialPos := ship.GetPosition()

	logging.Info("Distance: %v", d)

	travelAction := actions.NewMoveEntityAction(
		ship,
		dest,
		order.ExecuteTick,
		core.Tick(t),
	)

	returnAction := actions.NewMoveEntityAction(
		ship,
		models.Destination{Position: initialPos},
		order.ExecuteTick+core.Tick(t),
		core.Tick(t),
	)

	order.Actions = append(order.Actions, travelAction, returnAction)

	if dest.Entity != nil {
		scoutEntityAction := actions.NewScoutEntityAction(dest.Entity, order.ExecuteTick, core.Tick(t)*2)
		order.Actions = append(order.Actions, scoutEntityAction)
	}

	return order.Order

}
