package main

import (
	"fmt"
	"math/rand"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
)

const NUM_STAR_SYSTEMS = 3

func main() {

	// intialize systems
	for range NUM_STAR_SYSTEMS {
		system := GameStateGlobal.GenerateStarSystem()
		GameStateGlobal.StarSystems = append(GameStateGlobal.StarSystems, &system)
	}

	GameStateGlobal.Player = Player{Position{0, 0, 0}}
	logging.Ok("Player Initialized")

	GameStateGlobal.ShipManager.Ships = make(map[int]*Ship)

	for range 5 {
		name := fmt.Sprintf("Hermes %03d", rand.Intn(1000))
		ship := CreateNewShip(name, GameStateGlobal.Player.Position, Scout)

		GameStateGlobal.ShipManager.AddShip(ship)
	}

	logging.Ok("Ships Initialized")

	logging.Ok("State Initialized")
	engine.RunGame(&GameStateGlobal)
}
