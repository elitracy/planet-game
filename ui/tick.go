package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	TICKS_PER_SECOND_UI = 2
	TICK_SLEEP_UI       = time.Second / TICKS_PER_SECOND_UI
)

type tickMsg struct {
	Count int
}

func tick(count int) tea.Cmd {
	return tea.Tick(TICK_SLEEP_UI, func(t time.Time) tea.Msg { return tickMsg{Count: count + 1} })

}
