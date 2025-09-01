package systems

import (
	"github.com/elitracy/planets/models"
)

func TickPopulation(gs *models.GameState) {
	for _, starSystem := range gs.StarSystems {
		for _, planet := range starSystem.Planets {
			if gs.CurrentTick%5 == 0 {
				planet.Population += planet.PopulationGrowthRate
			}
			planet.Resources.Food.Quantity -= planet.Population * planet.Resources.Food.ConsumptionRate
		}
	}
}
