package systems

import (
	"github.com/elitracy/planets/core/state"
)

func TickPopulation() {
	for _, starSystem := range state.State.StarSystems {
		for _, planet := range starSystem.Planets {
			planet.Resources.Food.Quantity -= planet.Population * planet.Resources.Food.ConsumptionRate
		}
	}
}
