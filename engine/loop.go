package engine

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/logging"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
	"github.com/elitracy/planets/ui"
)

const TICK_SLEEP = time.Second

var PLAYER_START_LOC = models.Position{0, 0, 0}

func RunGame(state *models.GameState) {
	quit := make(chan struct{})

	origin := state.Player
	dest := state.StarSystems[0].Planets[0]

	travelStatus := ui.PaneManager.AddPane(ui.NewTravelStatusPane("Travel",
		1,
		origin.Position,
		dest.Position,
		"Player",
		dest.Name,
		state.CurrentTick,
		10,
		state,
	))
	loadingBar := ui.PaneManager.AddPane(ui.NewLoadingBarPane("Loading Bar"))
	planetList := ui.PaneManager.AddPane(ui.NewPlanetList(state.StarSystems[0].Planets, "Planet List"))

	grid := [][]int{
		{travelStatus, loadingBar},
		{loadingBar},
		{planetList},
	}
	dashboard := ui.PaneManager.AddPane(ui.NewDashboard(grid, 0, 0, "Dashboard"))

	ui.PaneManager.PushFocusStack(dashboard)

	p := tea.NewProgram(&ui.PaneManager)

	go func() {
		if _, err := p.Run(); err != nil {
			logging.Log(fmt.Sprintf("Alas, there's been an error: %v", err), "UI", "ERROR")
			os.Exit(1)
		}

		close(quit)
	}()

	logging.Log("Layout Initialized ✅", "LOOP")

	for {
		select {
		case <-quit:
			logging.Log("UI exited core loop", "UI")
			return
		default:
			// advance time
			state.CurrentTick++

			// update systems
			systems.TickConstructions(state)
			systems.TickStabilities(state)
			systems.TickPopulation(state)
			systems.TickPayloads(state)

			time.Sleep(TICK_SLEEP)

		}
	}

}
