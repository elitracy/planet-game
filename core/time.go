package core

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core/consts"
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

type UITickMsg struct {
	Tick
}

func UITickCmd(tick Tick) tea.Cmd {
	return tea.Tick(consts.TICK_SLEEP_UI, func(time.Time) tea.Msg { return UITickMsg{Tick: tick + 1} })
}

func TickCmd(tick Tick) tea.Cmd {
	return tea.Tick(consts.TICK_SLEEP, func(time.Time) tea.Msg { return TickMsg{Tick: tick + 1} })
}

type Pulse int64   // minutes
type Cycle float64 // years

func (t Tick) String() string {
	total := int64(t)

	ticks := total % int64(consts.TICKS_PER_PULSE)
	total /= int64(consts.TICKS_PER_PULSE)

	pulses := total % consts.PULSES_PER_CYCLE
	pulses = max(1, pulses)
	total /= pulses

	cycles := total / consts.PULSES_PER_CYCLE

	return fmt.Sprintf("%d.%03d.%04d", cycles, pulses, ticks)
}
