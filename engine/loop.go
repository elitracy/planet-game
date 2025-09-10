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
		{ui.NewTitlePane("Top Left", 0), ui.NewTitlePane("Top Right", 1)},
		{ui.NewTitlePane("Bottom Left", 2), ui.NewTitlePane("Bottom Right", 3)},
	}

	m := ui.Dashboard{
		Grid:      grid,
		ActiveRow: 0,
		ActiveCol: 0,
	}

	p := tea.NewProgram(m)

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
