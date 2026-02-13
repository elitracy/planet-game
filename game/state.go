package game

import (
	"math/rand"
	"slices"
	"sort"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/actions"
	"github.com/elitracy/planets/game/models"
	"github.com/elitracy/planets/game/orders"
)

var State *GameState

const (
	MIN_PLANETS = 2
	MAX_PLANETS = 10

	// center of system
	MIN_STAR_SYSTEM_DIST = 10_000
	MAX_STAR_SYSTEM_DIST = 1_000_000

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
	CurrentTick     engine.Tick
	StarSystems     []*models.StarSystem
	Player          models.Player
	OrderScheduler  engine.EventScheduler[*orders.Order]
	ActionScheduler engine.EventScheduler[*actions.Action]
	CompletedOrders []*orders.Order
	ShipManager     models.ShipManager
}

func (gs *GameState) CreatePlayer(location models.Location) models.Player {
	player := models.Player{Location: location}
	gs.Player = player
	return player
}

func (gs *GameState) GenerateStarSystem() *models.StarSystem {

	system_name_idx := rand.Intn(len(system_names))
	system_name := system_names[system_name_idx]

	system_names = slices.Delete(system_names, system_name_idx, system_name_idx+1)

	system_position := engine.Position{
		X: int64(rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST),
		Y: int64(rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST),
		Z: int64(rand.Intn(MAX_STAR_SYSTEM_DIST-MIN_STAR_SYSTEM_DIST) + MIN_STAR_SYSTEM_DIST),
	}

	system := models.CreateStarSystem(system_name, []*models.Planet{}, system_position)

	num_planets := rand.Intn(MAX_PLANETS-MIN_PLANETS) + MIN_PLANETS
	var planet_positions []engine.Position

	for range num_planets {
		position := engine.Position{
			X: int64(rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST)) + MIN_PLANET_DIST + system_position.X,
			Y: int64(rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST)) + MIN_PLANET_DIST + system_position.Y,
			Z: int64(rand.Intn(MAX_PLANET_DIST-MIN_PLANET_DIST)) + MIN_PLANET_DIST + system_position.Z,
		}
		planet_positions = append(planet_positions, position)
	}

	sort.Slice(planet_positions, func(i, j int) bool {
		d_i := engine.EuclidianDistance(planet_positions[i], system_position)
		d_j := engine.EuclidianDistance(planet_positions[j], system_position)
		return d_i < d_j
	})

	for i := range num_planets {
		starting_population := rand.Intn(MAX_START_POP-MIN_START_POP) + MIN_START_POP

		planet := models.CreatePlanet(
			system_name+"-"+planet_names[i],
			planet_positions[i],
			starting_population,
			STARTING_FARMS,
			STARTING_MINES,
			STARTING_SOLAR_GRIDS,
		)
		system.Planets = append(system.Planets, &planet)

	}

	return system

}

func (state *GameState) PushOrder(order *orders.Order) {
	state.OrderScheduler.Push(order)
	for _, action := range order.Actions {
		state.ActionScheduler.Push(action)
	}
}
