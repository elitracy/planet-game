package game

import (
	"math/rand"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/events"
	"github.com/elitracy/planets/game/models"
)

type EventChance struct {
	Probability float64
	New         func(models.Entity, engine.Tick) *events.Event
}

var EventTable = []EventChance{
	{0.1, events.NewPirateRaid},
}

func (state *GameState) TickEvents() {
	if State.CurrentTick%(config.TICKS_PER_PULSE) != 0 {
		return
	}

	for _, system := range state.StarSystems {
		for _, planet := range system.Planets {
			for _, chance := range EventTable {
				if planet.Colonized && rand.Float64() < chance.Probability {

				}

			}

		}
	}
}
