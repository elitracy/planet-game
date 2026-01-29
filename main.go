package main

import (
	"fmt"
	"math/rand"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/state"
	"github.com/elitracy/planets/ui"
)

const NUM_STAR_SYSTEMS = 3

const START_YEAR_TICK = 2049 * consts.TICKS_PER_CYCLE

func main() {

	state.State = &state.GameState{}
	state.State.CurrentTick = core.Tick(rand.Intn(START_YEAR_TICK) + START_YEAR_TICK)
	logging.SetTick(&state.State.CurrentTick)

	ui.InitPaneManager()

	for range NUM_STAR_SYSTEMS {
		system := state.State.GenerateStarSystem()
		state.State.StarSystems = append(state.State.StarSystems, system)
	}

	startingSystem := state.State.StarSystems[0]
	startingSystem.Colonized = true

	startingPlanet := startingSystem.Planets[0]

	for _, planet := range startingSystem.Planets {
		planet.Colonized = true
	}

	state.State.CreatePlayer(startingPlanet.GetLocation())
	logging.Ok("Player Initialized")

	state.State.ShipManager.Ships = make(map[int]*models.Ship)

	for range 5 {
		name := fmt.Sprintf("Hermes %03d", rand.Intn(1000))
		ship := models.CreateNewShip(name, startingPlanet.GetLocation(), models.Scout)

		state.State.ShipManager.AddShip(ship)
	}

	logging.Ok("Ships Initialized")

	logging.Ok("State Initialized")
	engine.RunGame()
}
