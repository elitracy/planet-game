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

	planetList := ui.NewPlanetListPane(state.State.StarSystems[0].Planets, "Planet List")
	orderList := ui.NewOrderStatusPane(&state.State.OrderScheduler, "Orders")
	systemsPane := ui.NewSystemsPane("Systems", state.State.StarSystems)

	ui.PaneManager.AddPane(planetList)
	ui.PaneManager.AddPane(orderList)
	ui.PaneManager.AddPane(systemsPane)

	ui.PaneManager.AddTab(planetList)
	ui.PaneManager.AddTab(orderList)
	ui.PaneManager.AddTab(systemsPane)

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
