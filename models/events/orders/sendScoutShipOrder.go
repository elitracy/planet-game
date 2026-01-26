package orders

import (
	"github.com/elitracy/planets/core"
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
	order := &ScoutPositionOrder{
		Order: &Order{
			Name:        "Scout destination",
			ExecuteTick: execTick,
			Status:      events.EventPending,
		},
		ship:        ship,
		destination: dest,
	}

	d := core.EuclidianDistance(ship.Position, dest.Position)
	t := d / ship.Velocity.Vector()

	initialPos := ship.GetPosition()

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

	if system, ok := dest.Entity.(*models.StarSystem); ok {

		scoutSystemAction := actions.NewScoutSystemAction(system, order.ExecuteTick, core.Tick(t)*2)
		order.Actions = append(order.Actions, scoutSystemAction)
	}

	return order.Order

}
