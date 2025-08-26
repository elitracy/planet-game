package systems

import "github.com/elitracy/planets/models"

func TickStabilities(gs *models.GameState) {
	for _, starSystem := range gs.StarSystems {
		for _, planet := range starSystem.Planets {

			if planet.Resources.Food.Quantity < 0 {
				planet.Stabilities.Unrest.Quantity++
			}

		}
	}
}
