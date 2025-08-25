package engine

import (
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
)

func RunGame(state *models.GameState) {

	for {
		// advance time
		state.CurrentTick++

		// update systems

		systems.TickConstructions(state)
		systems.TickStabilities(state)
		systems.TickPopulation(state)
		systems.TickPayloads(state)

		// render UI

		// handle Input

	}

}
