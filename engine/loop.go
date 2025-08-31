package engine

import (
	"fmt"
	"time"

	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
)

const TICK_SLEEP = 2 * time.Second

func RunGame(state *models.GameState) {

	for {
		// advance time
		state.CurrentTick++

		// update systems
		systems.TickConstructions(state)
		systems.TickStabilities(state)
		systems.TickPopulation(state)
		systems.TickPayloads(state)

		fmt.Println("\nTICK:", state.CurrentTick)

		// render UI

		// handle Input

		time.Sleep(TICK_SLEEP)
	}

}
