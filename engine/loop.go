package engine

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Game interface {
	Update(tick Tick)
}

func RunGame(g Game, model tea.Model, startTick Tick) {
	quit := make(chan struct{})

	p := tea.NewProgram(model, tea.WithAltScreen())

	go func() {
		if _, err := p.Run(); err != nil {
			os.Exit(1)
		}
		close(quit)
	}()

	tick := startTick
	for {
		select {
		case <-quit:
			return
		default:
			g.Update(tick)
			tick++
			time.Sleep(TICK_SLEEP)
		}
	}

}
