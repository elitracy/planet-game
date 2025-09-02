package engine

import (
	"fmt"
	"time"

	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
)

const TICK_SLEEP = 2 * time.Second

var PLAYER_START_LOC = models.Location{Coordinates: models.Coordinates{X: 0, Y: 0}}

func RunGame(state *models.GameState) {

	// player := state.CreatePlayer(PLAYER_START_LOC)

	for {
		// advance time [thread 1]
		state.CurrentTick++
		fmt.Println("TICK:", state.CurrentTick)

		// update systems
		systems.TickConstructions(state)
		systems.TickStabilities(state)
		systems.TickPopulation(state)
		systems.TickPayloads(state)

		// render UI [thread 2]

		// display unread messages
		// display planet options

		// handle Input [main thread]

		time.Sleep(TICK_SLEEP)
	}

}
