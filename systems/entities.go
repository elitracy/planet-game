package systems

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/state"
)

func TickSystems() {

	if state.State.CurrentTick%(core.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, starSystem := range state.State.StarSystems {
		starSystem.Tick()
	}

}
