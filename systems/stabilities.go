package systems

import (
	"github.com/elitracy/planets/core/state"
)

func TickStabilities() {
	for _, starSystem := range state.State.StarSystems {
		for _, planet := range starSystem.Planets {

			if planet.Resources.Food.Quantity < 0 {
				planet.Stabilities.Unrest.Quantity++
			}

		}
	}
}
