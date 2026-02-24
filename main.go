package main

import (
	"fmt"
	"math/rand"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/events"
	"github.com/elitracy/planets/game/models"
	"github.com/elitracy/planets/ui"
)

const NUM_STAR_SYSTEMS = 3
const START_YEAR_TICK = 2049 * config.TICKS_PER_CYCLE

func InitState() {
	game.State = &game.GameState{
		ColonizedPlanets: make(map[models.EntityID]*models.Planet),
		Ships:            make(map[models.EntityID]*models.Ship),
		CurrentTick:      engine.Tick(rand.Intn(START_YEAR_TICK) + START_YEAR_TICK),
		EventManager:     events.NewEventManager(),
		EntityManager:    models.NewEntityManager(),
	}

	engine.SetTick(&game.State.CurrentTick)

	for range NUM_STAR_SYSTEMS {
		system := game.State.GenerateStarSystem()
		game.State.StarSystems = append(game.State.StarSystems, system)
	}

	startingSystem := game.State.StarSystems[0]
	startingSystem.Colonized = true

	startingPlanet := startingSystem.Planets[0]

	for _, planet := range startingSystem.Planets {
		planet.Colonized = true
		game.State.ColonizedPlanets[planet.GetID()] = planet
	}

	player := models.NewPlayer(startingPlanet.GetLocation())

	game.State.Player = player
	game.State.EntityManager.Add(game.State.Player)
	engine.Ok("Player Initialized")

	for range 5 {
		name := fmt.Sprintf("Hermes %03d", rand.Intn(1000))
		ship := models.CreateNewShip(name, startingPlanet.GetLocation(), models.Scout)

		game.State.EntityManager.Add(ship)
		game.State.Ships[ship.GetID()] = ship
	}

	engine.Ok("State Initialized")
}

func InitUI() {
	ui.InitPaneManager()

	orderStatusList := ui.NewOrderStatusListPane("Orders", &game.State.OrderScheduler)
	systemsPane := ui.NewStarSystemListPane("Systems", game.State.StarSystems)

	ui.PaneManager.AddPane(orderStatusList)
	ui.PaneManager.AddPane(systemsPane)

	ui.PaneManager.AddTab(systemsPane)
	ui.PaneManager.AddTab(orderStatusList)

}

func main() {
	InitState()
	InitUI()

	engine.RunGame(game.State, ui.PaneManager, game.State.CurrentTick)
}
