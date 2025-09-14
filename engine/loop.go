package engine

import (
	"fmt"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/logging"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/systems"
	"github.com/elitracy/planets/ui"
)

const TICK_SLEEP = time.Second

var PLAYER_START_LOC = models.Location{Coordinates: models.Coordinates{X: 0, Y: 0}}

func RunGame(state *models.GameState) {
	quit := make(chan struct{})

	// render UI
	// p := tea.NewProgram(ui.CreatePlanetListInitialModel(state.StarSystems[0].Planets))

	grid := [][]tea.Model{
		{ui.NewCreateColonyPane("Create Colony", 1, state.StarSystems[0].Planets[0]), ui.NewTitlePane("Top Right", 2)},
		{ui.NewTitlePane("Bottom Left", 3), ui.NewTitlePane("Bottom Right", 4)},
	}

	m := ui.NewDashboard(grid, 0, 0, 0, "Dashboard")
	ui.PushFocus(&m)

	p := tea.NewProgram(&m)

	go func() {
		if _, err := p.Run(); err != nil {
			logging.Log(fmt.Sprintf("Alas, there's been an error: %v", err), "UI", "ERROR")
			os.Exit(1)
		}

		close(quit)
	}()

	for {
		select {
		case <-quit:
			logging.Log("UI exited core loop", "UI")
			return
		default:
			logging.Log(fmt.Sprintf("Active Pane: %d", ui.ActivePane().(ui.BasePane).GetId()), "LOOP")

			// advance time
			state.CurrentTick++
			logging.Log("TICK: "+strconv.Itoa(state.CurrentTick), "CORE")

			// update systems
			systems.TickConstructions(state)
			systems.TickStabilities(state)
			systems.TickPopulation(state)
			systems.TickPayloads(state)

			time.Sleep(TICK_SLEEP)

		}
	}

}
