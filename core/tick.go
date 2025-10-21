package core

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	TICKS_PER_SECOND_UI = 2
	TICK_SLEEP_UI       = time.Second / TICKS_PER_SECOND_UI
)

type Tick int64

func (t Tick) ToDuration(tickRate int) time.Duration {
	return time.Duration(t) * time.Second / time.Duration(tickRate)
}

func FromDuration(d time.Duration, tickRate int) Tick {
	return Tick(int64(d.Seconds() * float64(tickRate)))
}

type TickMsg struct {
	Tick
}

func TickCmd(tick Tick) tea.Cmd {
	return tea.Tick(TICK_SLEEP_UI, func(time.Time) tea.Msg { return TickMsg{Tick: tick + 1} })

}
