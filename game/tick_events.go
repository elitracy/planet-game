package game

import (
	"math/rand"

	"github.com/elitracy/planets/game/models"
)

func (state *GameState) TickEvents() {
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

	if state.CurrentTick < state.EventManager.NextEvent {
		return
	}

	var planet *models.Planet
	planetIndex := rand.Intn(len(state.ColonizedPlanets))
	for _, p := range state.ColonizedPlanets {
		if planetIndex == 0 {
			planet = p
			break
		}
		planetIndex--
	}

	for _, chance := range state.EventManager.Events {
		if rand.Float64() < chance.Probability {
			event := chance.New(planet, state.CurrentTick)

			state.EventManager.Add(event, state.CurrentTick)
			state.EventManager.RollNextEventInterval(state.CurrentTick)
			break
		}
	}

}
