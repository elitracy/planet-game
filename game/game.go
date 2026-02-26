package game

import (
	"github.com/elitracy/planets/engine"
)

func (state *GameState) Update(tick engine.Tick) {
	state.CurrentTick = tick

	state.TickEvents()
	state.TickTasks()
	state.TickSystems()
}
