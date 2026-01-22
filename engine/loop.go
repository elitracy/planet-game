package engine

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
	"github.com/elitracy/planets/systems"
	"github.com/elitracy/planets/ui"
)

var PLAYER_START_LOC = core.Position{X: 0, Y: 0, Z: 0}

func RunGame() {
	quit := make(chan struct{})

	orderList := ui.NewOrderStatusPane(&state.State.OrderScheduler, "Orders")
	systemsPane := ui.NewSystemListPane("Systems", state.State.StarSystems)

	ui.PaneManager.AddPane(orderList)
	ui.PaneManager.AddPane(systemsPane)

	ui.PaneManager.AddTab(systemsPane)
	ui.PaneManager.AddTab(orderList)

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
			systems.TickSystems()

			time.Sleep(core.TICK_SLEEP)
		}
	}

}
