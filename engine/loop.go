package engine

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
	"github.com/elitracy/planets/ui"
)

var PLAYER_START_LOC = core.Position{X: 0, Y: 0, Z: 0}

func RunGame() {
	quit := make(chan struct{})

	planetList := ui.PaneManager.AddPane(ui.NewPlanetListPane(state.State.StarSystems[0].Planets, "Planet List"))
	orderList := ui.PaneManager.AddPane(ui.NewOrderStatusPane(&state.State.OrderScheduler, "Orders"))
	tabSelect := ui.PaneManager.AddPane(ui.NewTabSelectPane("Tabs", []int{planetList, orderList}))

	grid := [][]int{
		{tabSelect, planetList},
	}
	dashboard := ui.PaneManager.AddPane(ui.NewDashboard(grid, 0, 0, "Dashboard"))

	ui.PaneManager.PushFocusStack(dashboard)

	p := tea.NewProgram(ui.PaneManager, tea.WithAltScreen())

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
			state.State.Tick++

			systems.TickOrderScheduler()
			systems.TickActionScheduler()
			systems.TickConstructions()
			systems.TickStabilities()
			systems.TickPopulation()

			time.Sleep(TICK_SLEEP)
		}
	}

}
