package systems

import (
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/state"
)

func TickSystems() {

	if state.State.CurrentTick%(consts.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, starSystem := range state.State.StarSystems {
		starSystem.Tick()
	}

}
