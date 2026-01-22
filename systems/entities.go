package systems

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
)

func TickSystems() {

	if state.State.Tick%(core.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, starSystem := range state.State.StarSystems {
		starSystem.Tick()
	}

}
