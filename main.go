package main

import (
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

	logging.Ok("State Initialized")

	GameStateGlobal.Player = Player{Position{0, 0, 0}}
	GameStateGlobal.ShipManager.CreateShip("Hermes I", GameStateGlobal.Player.Position, Scout)

	logging.Ok("Player Initialized")

	engine.RunGame(&GameStateGlobal)
}
