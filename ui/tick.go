package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/logging"
)

type tickMsg struct {
	Count int
}

func tick(count int) tea.Cmd {
	logging.Log(fmt.Sprintf("Count: %d", count), "TICK")
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg{Count: count + 1} })

}
