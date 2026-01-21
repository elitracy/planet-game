package systems

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
)

const CONSTRUCTION_SLEEP = core.TICKS_PER_SECOND * 2

func TickConstructions() {

	if state.State.Tick%CONSTRUCTION_SLEEP != 0 {
		return
	}

	for _, starSystem := range state.State.StarSystems {
		for _, planet := range starSystem.Planets {

			for _, farm := range planet.Constructions.Farms {
				planet.Resources.Food.Quantity += farm.Quantity * farm.ProductionRate
			}

			for _, mine := range planet.Constructions.Mines {
				planet.Resources.Minerals.Quantity += mine.Quantity * mine.ProductionRate
			}

			for _, solarGrid := range planet.Constructions.SolarGrids {
				planet.Resources.Energy.Quantity += solarGrid.Quantity * solarGrid.ProductionRate
			}
		}
	}
}
