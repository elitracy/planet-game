package systems

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
)

const FOOD_PER_PERSON = 3
const POPULATION_GROWTH_RATE = 1000

func TickPopulation() {

	if state.State.Tick%(core.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, starSystem := range state.State.StarSystems {
		for _, planet := range starSystem.Planets {

			currentFood := planet.Resources.Food.Quantity
			requiredFood := planet.Population * FOOD_PER_PERSON

			if currentFood < requiredFood {
				planet.Stabilities.Happiness.Quantity--

			} else {
				planet.Resources.Food.Quantity -= requiredFood

				currentFood = planet.Resources.Food.Quantity
				addPopRequiredFood := FOOD_PER_PERSON * POPULATION_GROWTH_RATE

				if currentFood >= addPopRequiredFood {
					planet.Resources.Food.Quantity -= FOOD_PER_PERSON * POPULATION_GROWTH_RATE
					planet.Population += POPULATION_GROWTH_RATE
				}
			}

		}
	}

}
