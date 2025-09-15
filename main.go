package main

import (
	// "fmt"

	"github.com/elitracy/planets/engine"
	// "github.com/elitracy/planets/logging"
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

	gameState.Player = Player{Position{0, 0, 0}}

	// logging.Log(fmt.Sprint(gameState.StarSystems), "MAIN")

	engine.RunGame(&gameState)
}
