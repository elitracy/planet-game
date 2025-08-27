package systems

import (
	"fmt"

	"github.com/elitracy/planets/models"
)

func TickConstructions(gs *models.GameState) {
	for _, starSystem := range gs.StarSystems {
		for _, planet := range starSystem.Planets {

			food_acc := 0
			for _, farm := range planet.Constructions.Farms {
				planet.Resources.Food.Quantity += farm.Quantity
				food_acc += farm.Quantity
			}
			fmt.Printf("Food Generated: %d\n", food_acc)

			minerals_acc := 0
			for _, mine := range planet.Constructions.Mines {
				planet.Resources.Minerals.Quantity += mine.Quantity
				minerals_acc += mine.Quantity
			}

			fmt.Printf("Minerals Generated: %d\n", minerals_acc)

			energy_acc := 0
			for _, solarGrid := range planet.Constructions.SolarGrids {
				planet.Resources.Energy.Quantity += solarGrid.Quantity
				energy_acc += solarGrid.Quantity
			}
			fmt.Printf("Energy Generated: %d\n", energy_acc)
		}
	}
}
