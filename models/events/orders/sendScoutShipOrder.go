package orders

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/actions"
)

type ScoutPositionOrder struct {
	*Order
	ship        *models.Ship
	destination models.Location
}

func NewScoutDestinationOrder(ship *models.Ship, dest models.Location, startTick core.Tick) *Order {
	name := ""
	if dest.Entity != nil {
		name = dest.Entity.GetName()
	} else {
		name = dest.Entity.GetLocation().Position.String()
	}
	d := core.EuclidianDistance(ship.GetLocation().Position, dest.Position)
	t := d / ship.Velocity.Vector()

	initialLocation := ship.Location

	order := &ScoutPositionOrder{
		Order: &Order{
			Name:      "Scout " + name,
			StartTick: startTick,
			Status:    events.EventPending,
		},
		ship:        ship,
		destination: dest,
	}

	travelAction := actions.NewMoveEntityAction(
		ship,
		dest,
		order.StartTick,
		core.Tick(t),
	)
	order.AddAction(travelAction.Action)

	returnActionStartTick := travelAction.GetStartTick() + travelAction.GetDuration()

	if dest.Entity != nil {
		scoutEntityAction := actions.NewScoutEntityAction(dest.Entity, travelAction.GetStartTick()+travelAction.GetDuration(), consts.TICKS_PER_PULSE*2)
		order.AddAction(scoutEntityAction.Action)

		returnActionStartTick += scoutEntityAction.GetDuration()
	}

	returnAction := actions.NewMoveEntityAction(
		ship,
		initialLocation,
		returnActionStartTick,
		core.Tick(t),
	)
	returnAction.Description = fmt.Sprintf("Return %v to %v", returnAction.TargetEntity.GetName(), initialLocation)

	order.AddAction(returnAction.Action)

	return order.Order

}
