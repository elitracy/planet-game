package systems

import "github.com/elitracy/planets/core/state"

func TickConstructions() {
	for _, starSystem := range state.State.StarSystems {
		for _, planet := range starSystem.Planets {

			food_acc := 0
			for _, farm := range planet.Constructions.Farms {
				planet.Resources.Food.Quantity += farm.Quantity
				food_acc += farm.Quantity
			}

			minerals_acc := 0
			for _, mine := range planet.Constructions.Mines {
				planet.Resources.Minerals.Quantity += mine.Quantity
				minerals_acc += mine.Quantity
			}

			energy_acc := 0
			for _, solarGrid := range planet.Constructions.SolarGrids {
				planet.Resources.Energy.Quantity += solarGrid.Quantity
				energy_acc += solarGrid.Quantity
			}
		}
	}
}
