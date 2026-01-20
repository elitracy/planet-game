package core

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Tick int64

const (
	TICKS_PER_SECOND_UI = 2
	TICK_SLEEP_UI       = time.Second / TICKS_PER_SECOND_UI
)


func (t Tick) ToDuration(tickRate int) time.Duration {
	return time.Duration(t) * time.Second / time.Duration(tickRate)
}

func FromDuration(d time.Duration, tickRate int) Tick {
	return Tick(int64(d.Seconds() * float64(tickRate)))
}

type TickMsg struct {
	Tick
}

type UITickMsg struct {
	Tick
}

func UITickCmd(tick Tick) tea.Cmd {
	return tea.Tick(TICK_SLEEP_UI, func(time.Time) tea.Msg { return UITickMsg{Tick: tick + 1} })
}

func TickCmd(tick Tick) tea.Cmd {
	return func() tea.Msg { return TickMsg{Tick: tick + 1} }}



// Megacycle = 1 year
// Kiolcycle = 1 month
// Cycle = 1 day = 100,000 ticks
// Milicycle = 1 minute
// Centicycle = 1 minute
type Cycle float64

func TickToCycle(t Tick) Cycle {
	return Cycle(float64(t) / float64(100_000))
}

func CycleToTick(c Cycle) Tick {
	return Tick(c * 100_000)
}
