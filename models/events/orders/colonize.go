package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/actions"
)

type CreateColonyOrder struct {
	*Order
	planet *models.Planet
}

func NewCreateColonyOrder(planet *models.Planet, startTick core.Tick) *Order {
	order := &CreateColonyOrder{
		Order: &Order{
			Name:      "Create Colony",
			StartTick: startTick,
			Status:    events.EventPending,
		},
		planet: planet,
	}

	createFarmAction := actions.NewBuildFarmAction(
		planet,
		order.StartTick,
		consts.TICKS_PER_SECOND*10,
	)

	createMineAction := actions.NewBuildMineAction(
		planet,
		order.StartTick,
		consts.TICKS_PER_SECOND*10,
	)

	createSolarGridAction := actions.NewBuildSolarGridAction(
		planet,
		order.StartTick,
		consts.TICKS_PER_SECOND*10,
	)

	colonizeAction := actions.NewColonizeAction(planet, order.StartTick, consts.TICKS_PER_SECOND)

	order.Actions = append(order.Actions, createFarmAction.Action, createMineAction.Action, createSolarGridAction.Action, colonizeAction.Action)

	return order.Order
}
