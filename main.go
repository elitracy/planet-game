package main

import (
	"github.com/elitracy/planets/engine"
	. "github.com/elitracy/planets/models"
)

const STARTING_POPULATION = 1000
const STARTING_POPULATION_GROWTH_RATE = 100

const STARTING_FARMS = 2
const STARTING_MINES = 2
const STARTING_SOLAR_GRIDS = 2

const STARTING_FOOD = 5000
const STARTING_FOOD_CONSUMPTION_RATE = 1
const STARTING_MINERAL = 5000
const STARTING_MINERAL_CONSUMPTION_RATE = 1
const STARTING_ENERGY = 5000
const STARTING_ENERGY_CONSUMPTION_RATE = 1

func main() {

	gameState := GameState{}

	systemA := &StarSystem{}
	systemB := &StarSystem{}

	earth := CreatePlanet("EARTH", 0, 0, STARTING_POPULATION, STARTING_POPULATION_GROWTH_RATE, STARTING_FOOD, STARTING_MINERAL, STARTING_ENERGY, STARTING_FOOD_CONSUMPTION_RATE, STARTING_MINERAL_CONSUMPTION_RATE, STARTING_ENERGY_CONSUMPTION_RATE, STARTING_FARMS, STARTING_MINES, STARTING_SOLAR_GRIDS)
	james := CreatePlanet("JAMES", 0, 0, STARTING_POPULATION, STARTING_POPULATION_GROWTH_RATE, STARTING_FOOD, STARTING_MINERAL, STARTING_ENERGY, STARTING_FOOD_CONSUMPTION_RATE, STARTING_MINERAL_CONSUMPTION_RATE, STARTING_ENERGY_CONSUMPTION_RATE, STARTING_FARMS, STARTING_MINES, STARTING_SOLAR_GRIDS)

	systemA.Planets = append(systemA.Planets, &earth)
	systemB.Planets = append(systemB.Planets, &james)

	gameState.StarSystems = append(gameState.StarSystems, systemA, systemB)

	player2 := Player{Location: james.Location}
	player2.SendMessagePayload("Hello from James", earth, gameState.CurrentTick)

	player1 := Player{Location: earth.Location}
	player1.SendMessagePayload("Hello from Earth", james, gameState.CurrentTick)

	engine.RunGame(&gameState)
}
