package main

import (
	"fmt"
	"math/rand"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
	. "github.com/elitracy/planets/core/state"
	"github.com/elitracy/planets/engine"
	. "github.com/elitracy/planets/models"
)

const NUM_STAR_SYSTEMS = 3

func main() {

	// intialize systems
	for range NUM_STAR_SYSTEMS {
		system := State.GenerateStarSystem()
		State.StarSystems = append(State.StarSystems, &system)
	}

	State.Player = Player{core.Position{0, 0, 0}}
	logging.Ok("Player Initialized")

	State.ShipManager.Ships = make(map[int]*Ship)

	for range 5 {
		name := fmt.Sprintf("Hermes %03d", rand.Intn(1000))
		ship := CreateNewShip(name, State.Player.Position, Scout)

		State.ShipManager.AddShip(ship)
	}

	logging.Ok("Ships Initialized")

	logging.Ok("State Initialized")
	engine.RunGame()
}
