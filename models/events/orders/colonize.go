package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events/actions"
)

type CreateColonyOrder struct {
	*Order
	planet *models.Planet
}

func NewCreateColonyOrder(planet *models.Planet, execTick core.Tick) *Order {
	order := &CreateColonyOrder{
		Order: &Order{
			Name:        "Create Colony",
			ExecuteTick: execTick,
			Status:      consts.EventPending,
		},
		planet: planet,
	}

	createFarmAction := actions.NewBuildFarmAction(
		planet,
		order.ExecuteTick,
		core.TICKS_PER_SECOND*10,
	)

	createMineAction := actions.NewBuildMineAction(
		planet,
		order.ExecuteTick,
		core.TICKS_PER_SECOND*10,
	)

	createSolarGridAction := actions.NewBuildSolarGridAction(
		planet,
		order.ExecuteTick,
		core.TICKS_PER_SECOND*10,
	)

	colonizeAction := actions.NewColonizeAction(planet, order.ExecuteTick, core.TICKS_PER_SECOND)

	order.Actions = append(order.Actions, createFarmAction, createMineAction, createSolarGridAction, colonizeAction)

	return order.Order
}
