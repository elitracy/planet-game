package engine

import (
	"strconv"
	"time"

	"github.com/elitracy/planets/logging"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
)

const TICK_SLEEP = 2 * time.Second

var PLAYER_START_LOC = models.Location{Coordinates: models.Coordinates{X: 0, Y: 0}}

func RunGame(state *models.GameState) {
	for {
		// advance time
		state.CurrentTick++
		logging.Log("TICK: "+strconv.Itoa(state.CurrentTick), "CORE")

		// update systems
		systems.TickConstructions(state)
		systems.TickStabilities(state)
		systems.TickPopulation(state)
		systems.TickPayloads(state)

		// render UI

		time.Sleep(TICK_SLEEP)
	}

}
