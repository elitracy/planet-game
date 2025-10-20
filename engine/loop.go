package engine

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
	"github.com/elitracy/planets/ui"
)

var PLAYER_START_LOC = Position{X: 0, Y: 0, Z: 0}

func RunGame(state *GameState) {
	quit := make(chan struct{})

	planetList := ui.PaneManager.AddPane(ui.NewPlanetList(state.StarSystems[0].Planets, "Planet List"))
	orderList := ui.PaneManager.AddPane(ui.NewOrderStatusPane(&state.OrderScheduler, "Orders"))
	shipManager := ui.PaneManager.AddPane(ui.NewShipManagementPane("Ship Manager", &state.ShipManager, nil))

	grid := [][]int{
		{planetList, shipManager, orderList},
	}
	dashboard := ui.PaneManager.AddPane(ui.NewDashboard(grid, 0, 0, "Dashboard"))

	ui.PaneManager.PushFocusStack(dashboard)

	p := tea.NewProgram(&ui.PaneManager)

	go func() {
		if _, err := p.Run(); err != nil {
			os.Exit(1)
		}

		close(quit)
	}()

	for {
		select {
		case <-quit:
			return
		default:
			state.Tick++

			systems.TickOrderScheduler()
			systems.TickActionScheduler()
			systems.TickConstructions(state)
			systems.TickStabilities(state)
			systems.TickPopulation(state)

			time.Sleep(TICK_SLEEP)
		}
	}

}
