package systems

import "github.com/elitracy/planets/core/state"

func TickConstructions() {
	for _, starSystem := range state.State.StarSystems {
		for _, planet := range starSystem.Planets {

			for _, farm := range planet.Constructions.Farms {
				planet.Resources.Food.Quantity += farm.Quantity
			}

			for _, mine := range planet.Constructions.Mines {
				planet.Resources.Minerals.Quantity += mine.Quantity
			}

			for _, solarGrid := range planet.Constructions.SolarGrids {
				planet.Resources.Energy.Quantity += solarGrid.Quantity
			}
		}
	}
}
