package systems

import (
	"github.com/elitracy/planets/models"
)

func TickPopulation(gs *models.GameState) {
	for _, starSystem := range gs.StarSystems {
		for _, planet := range starSystem.Planets {
			planet.Resources.Food.Quantity -= planet.Population * planet.Resources.Food.ConsumptionRate
		}
	}
}
