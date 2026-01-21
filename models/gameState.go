package models

import (
	"math/rand"
	"slices"

	"github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core"
)

const (
	MIN_PLANETS = 2
	MAX_PLANETS = 10

	// center of system
	MIN_STAR_SYSTEM_DIST = 10000
	MAX_STAR_SYSTEM_DIST = 1000000

	// from star system center
	MIN_PLANET_DIST = 1000
	MAX_PLANET_DIST = 5000

	MIN_START_POP = 1000
	MAX_START_POP = 1_000_000

	STARTING_FARMS       = 2
	STARTING_MINES       = 2
	STARTING_SOLAR_GRIDS = 2
)

var (
	system_names     = []string{"Delta", "Zenith", "Umbra", "Roche", "Lagrange", "Hohmann", "Horizon", "Oberth", "Parallax", "Aphelion"}
	planet_names     = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	moon_names       = []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
	world_type_names = []string{"Colony", "Station", "Outpost", "Relay", "Belt", "Gate"}
	scout_ship_names = []string{"Hermes"}
)

type GameState struct {
	Tick            core.Tick
	StarSystems     []*StarSystem
	Player          Player
	OrderScheduler  EventScheduler[Order]
	ActionScheduler EventScheduler[Action]
	CompletedOrders []Order
	ShipManager
}

func (gs *GameState) CreatePlayer(location Position) Player {
	player := Player{Position: location}
	gs.Player = player
	return player
}

func (gs *GameState) GenerateStarSystem() StarSystem {

	system_name_idx := rand.Intn(len(system_names))
	system_name := system_names[system_name_idx]

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
			STARTING_FARMS,
			STARTING_MINES,
			STARTING_SOLAR_GRIDS,
		)
		system.Planets = append(system.Planets, &planet)
	}

	return system

}
