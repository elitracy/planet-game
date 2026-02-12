package game

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/engine"
)

type UITickMsg struct {
	engine.Tick
}

func UITickCmd(tick engine.Tick) tea.Cmd {
	return tea.Tick(TICK_SLEEP_UI, func(time.Time) tea.Msg { return UITickMsg{Tick: tick + 1} })
}

type Pulse int64   // minutes
type Cycle float64 // years

func TickPulseString(t engine.Tick) string {
	total := int64(t)

	ticks := total % int64(TICKS_PER_PULSE)
	total /= int64(TICKS_PER_PULSE)

	pulses := total % PULSES_PER_CYCLE
	pulses = max(1, pulses)
	total /= pulses

	cycles := total / PULSES_PER_CYCLE

	return fmt.Sprintf("%d.%03d.%04d", cycles, pulses, ticks)
}
