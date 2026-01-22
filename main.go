package main

import (
	"fmt"
	"math/rand"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
	state "github.com/elitracy/planets/core/state"
	"github.com/elitracy/planets/engine"
	models "github.com/elitracy/planets/models"
)

const NUM_STAR_SYSTEMS = 3

const START_YEAR_TICK = 2049 * core.TICKS_PER_CYCLE

func main() {

	state.State.Tick = core.Tick(rand.Intn(START_YEAR_TICK) + START_YEAR_TICK)

	for range NUM_STAR_SYSTEMS {
		system := state.State.GenerateStarSystem()
		state.State.StarSystems = append(state.State.StarSystems, &system)
	}

	startingSystem := state.State.StarSystems[0]
	startingSystem.Colonized = true

	startingPlanet := startingSystem.Planets[0]

	state.State.CreatePlayer(startingPlanet.Position)
	logging.Ok("Player Initialized")

	state.State.ShipManager.Ships = make(map[int]*models.Ship)

	for range 5 {
		name := fmt.Sprintf("Hermes %03d", rand.Intn(1000))
		ship := models.CreateNewShip(name, state.State.Player.Position, models.Scout)

		state.State.ShipManager.AddShip(ship)
	}

	logging.Ok("Ships Initialized")

	logging.Ok("State Initialized")
	engine.RunGame()
}
