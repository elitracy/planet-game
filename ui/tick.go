package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct {
	Count int
}

func tick(count int) tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg{Count: count + 1} })

}
