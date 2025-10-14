package models

import (
	"math/rand"
	"slices"
	"time"
)

var system_names = []string{"Delta", "Zenith", "Umbra", "Roche", "Lagrange", "Hohmann", "Horizon", "Oberth", "Parallax", "Aphelion"}
var planet_names = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var moon_names = []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
var world_type_names = []string{"Colony", "Station", "Outpost", "Relay", "Belt", "Gate"}

const MIN_PLANETS = 2
const MAX_PLANETS = 10

// center of system
const MIN_STAR_SYSTEM_DIST = 100
const MAX_STAR_SYSTEM_DIST = 1000

// from star system center
const MIN_PLANET_DIST = 10
const MAX_PLANET_DIST = 50

const MIN_START_POP = 1000
const MAX_START_POP = 1_000_000
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

const TICKS_PER_SECOND = 8
const TICK_SLEEP = time.Second / TICKS_PER_SECOND

type GameState struct {
	CurrentTick     int
	StarSystems     []*StarSystem
	Player          Player
	OrderScheduler  EventScheduler[*Order]
	ActionScheduler EventScheduler[*Action]
	CompletedOrders []*Order
	ShipManager
}

var GameStateGlobal GameState

func (gs *GameState) CreatePlayer(location Position) Player {
	player := Player{Position: location}
	return player
}

func (gs *GameState) GenerateStarSystem() StarSystem {

	system_name_idx := rand.Intn(len(system_names))
	system_name := system_names[system_name_idx]

	// remove system name from list
	system_names = slices.Delete(system_names, system_name_idx, system_name_idx+1)

	system_location := Position{
		X: rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST,
		Y: rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST,
		Z: rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST,
	}

	system := StarSystem{Name: system_name, Planets: []*Planet{}, Position: system_location}

	num_planets := rand.Intn(MAX_PLANETS-MIN_PLANETS) + MIN_PLANETS
	for i := range num_planets {

		starting_population := rand.Intn(MAX_START_POP-MIN_START_POP) + MIN_START_POP

		planet_location := Position{
			X: rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST) + MIN_PLANET_DIST,
			Y: rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST) + MIN_PLANET_DIST,
			Z: rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST) + MIN_PLANET_DIST,
		}

		planet := CreatePlanet(
			system_name+"-"+planet_names[i],
			planet_location.X,
			planet_location.Y,
			planet_location.Z,
			starting_population,
			STARTING_FOOD,
			STARTING_MINERAL,
			STARTING_ENERGY,
			STARTING_FOOD_CONSUMPTION_RATE,
			STARTING_MINERAL_CONSUMPTION_RATE,
			STARTING_ENERGY_CONSUMPTION_RATE,
			STARTING_FARMS,
			STARTING_MINES,
			STARTING_SOLAR_GRIDS,
		)
		system.Planets = append(system.Planets, &planet)
	}

	return system

}
