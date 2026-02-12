package systems

import "github.com/elitracy/planets/game"

func TickSystems() {

	if game.State.CurrentTick%(game.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, starSystem := range game.State.StarSystems {
		starSystem.Tick()
	}

}
