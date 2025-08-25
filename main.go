package main

import (
	"github.com/elitracy/planets/engine"
	. "github.com/elitracy/planets/models"
)


func main() {
	
	gameState := GameState{}

	engine.RunGame(&gameState)
}
