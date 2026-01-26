package state

import (
	"math/rand"
	"slices"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/actions"
	"github.com/elitracy/planets/models/events/orders"
)

var State *GameState

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
	CurrentTick            core.Tick
	StarSystems     []*models.StarSystem
	Player          models.Player
	OrderScheduler  events.EventScheduler[*orders.Order]
	ActionScheduler events.EventScheduler[*actions.Action]
	CompletedOrders []*orders.Order
	ShipManager     models.ShipManager
}

func (gs *GameState) CreatePlayer(location core.Position) models.Player {
	player := models.Player{Position: location}
	gs.Player = player
	return player
}

func (gs *GameState) GenerateStarSystem() *models.StarSystem {

	system_name_idx := rand.Intn(len(system_names))
	system_name := system_names[system_name_idx]

	system_names = slices.Delete(system_names, system_name_idx, system_name_idx+1)

	system_position := core.Position{
		X: rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST,
		Y: rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST,
		Z: rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST,
	}

	system := models.CreateStarSystem(system_name, []*models.Planet{}, system_position)

	num_planets := rand.Intn(MAX_PLANETS-MIN_PLANETS) + MIN_PLANETS
	for i := range num_planets {

		starting_population := rand.Intn(MAX_START_POP-MIN_START_POP) + MIN_START_POP

		planet_location := core.Position{
			X: rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST) + MIN_PLANET_DIST,
			Y: rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST) + MIN_PLANET_DIST,
			Z: rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST) + MIN_PLANET_DIST,
		}

		planet := models.CreatePlanet(
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
