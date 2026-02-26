package events

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/models"
)

func NewPirateRaid(targetEntity models.Entity, tick engine.Tick) *Event {
	return &Event{
		Name:        "Pirate Raid",
		Description: fmt.Sprintf("Pirates are attacking %v", targetEntity.GetName()),
		Severity:    Moderate,
		Target:      targetEntity,
		Resolved:    false,
		Effect: func() {
			if planet, ok := targetEntity.(*models.Planet); ok {
				planet.Resources.Minerals.Quantity -= int(float64(planet.Resources.Minerals.Quantity) * .1)
				planet.Resources.Food.Quantity -= int(float64(planet.Resources.Food.Quantity) * .1)
				planet.Resources.Energy.Quantity -= int(float64(planet.Resources.Energy.Quantity) * .1)
			}
		},
		Expire: func() {
			if planet, ok := targetEntity.(*models.Planet); ok {
				planet.Resources.Minerals.Quantity -= int(float64(planet.Resources.Minerals.Quantity) * .20)
				planet.Resources.Food.Quantity -= int(float64(planet.Resources.Food.Quantity) * .20)
				planet.Resources.Energy.Quantity -= int(float64(planet.Resources.Energy.Quantity) * .20)
			}
		},
		Duration: config.TICKS_PER_PULSE * 20,
		Interval: config.TICKS_PER_PULSE,
	}
}
