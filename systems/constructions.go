package systems

import (
	"github.com/elitracy/planets/models"
)

func TickConstructions(gs *models.GameState) {
	for _, starSystem := range gs.StarSystems {
		for _, planet := range starSystem.Planets {

			for _, farm := range planet.Constructions.Farm {
				planet.Resources.Food.Quantity += farm.Quantity
			}

			for _, mine := range planet.Constructions.Mine {
				planet.Resources.Minerals.Quantity += mine.Quantity
			}

			for _, solarGrid := range planet.Constructions.SolarGrid {
				planet.Resources.Energy.Quantity += solarGrid.Quantity
			}
		}
	}
}
