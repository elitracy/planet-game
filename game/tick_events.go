package game

import (
	"math/rand"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
)

func (state *GameState) TickEvents() {
	if State.CurrentTick%(config.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, system := range state.StarSystems {
		for _, planet := range system.Planets {
			for _, chance := range state.EventManager.Events {
				if planet.Colonized && rand.Float64() < chance.Probability {
					event := chance.New(planet, state.CurrentTick)
					engine.Info("Starting Event: %v: %v", event.Description, event.Target.GetName())
					state.EventManager.Add(event, state.CurrentTick)
				}
			}
		}
	}

	for _, event := range state.EventManager.ActiveEvents {
		elapsed := state.CurrentTick - event.Start

		if event.Resolved {
			continue
		}

		if elapsed >= event.Duration {
			event.Expire()
			event.Resolved = true
		}

		if event.Duration > 0 && elapsed < event.Duration && elapsed%event.Interval == 0 {
			event.Effect()
		}

	}
}
