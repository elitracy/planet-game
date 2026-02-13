package game

import "github.com/elitracy/planets/game/config"

func (state *GameState)TickSystems() {

	if State.CurrentTick%(config.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, starSystem := range state.StarSystems {
		starSystem.Tick()
	}

}
