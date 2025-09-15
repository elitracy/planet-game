package main

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
)

const NUM_STAR_SYSTEMS = 3

func main() {

	gameState := GameState{}

	// intialize systems
	for range NUM_STAR_SYSTEMS {
		system := gameState.GenerateStarSystem()
		gameState.StarSystems = append(gameState.StarSystems, &system)
	}

	logging.Log("State Initialized ✅", "MAIN")

	gameState.Player = Player{Position{0, 0, 0}}

	logging.Log("Player Initialized ✅", "MAIN")

	engine.RunGame(&gameState)
}
