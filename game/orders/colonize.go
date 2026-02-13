package orders

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/actions"
	"github.com/elitracy/planets/game/models"
)

type CreateColonyOrder struct {
	*Order
	planet *models.Planet
}

func NewCreateColonyOrder(planet *models.Planet, startTick engine.Tick) *Order {
	order := &CreateColonyOrder{
		Order: &Order{
			Name:      "Create Colony",
			StartTick: startTick,
			Status:    engine.EventPending,
		},
		planet: planet,
	}

	createFarmAction := actions.NewBuildFarmAction(
		planet,
		order.StartTick,
		engine.TICKS_PER_SECOND*10,
	)

	createMineAction := actions.NewBuildMineAction(
		planet,
		order.StartTick,
		engine.TICKS_PER_SECOND*10,
	)

	createSolarGridAction := actions.NewBuildSolarGridAction(
		planet,
		order.StartTick,
		engine.TICKS_PER_SECOND*10,
	)

	colonizeAction := actions.NewColonizeAction(
		planet,
		order.StartTick,
		engine.TICKS_PER_SECOND*10,
	)

	order.AddAction(createFarmAction.Action)
	order.AddAction(createMineAction.Action)
	order.AddAction(createSolarGridAction.Action)
	order.AddAction(colonizeAction.Action)

	return order.Order
}
