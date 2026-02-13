package events

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/models"
)

func NewPirateRaid(targetEntity models.Entity, tick engine.Tick) *Event {
	return &Event{
		Name:        "Pirate Raid",
		Description: fmt.Sprintf("Pirates are attacking %v", targetEntity.GetName()),
		Severity:    Moderate,
		Target:      targetEntity,
		Deadline:    tick + engine.TICKS_PER_SECOND*20,
		Resolved:    false,
		Effect: func() {
			if planet, ok := targetEntity.(*models.Planet); ok {
				planet.Resources.Minerals.Quantity -= int(float64(planet.Resources.Minerals.Quantity) * .5)
				planet.Resources.Food.Quantity -= int(float64(planet.Resources.Food.Quantity) * .5)
				planet.Resources.Energy.Quantity -= int(float64(planet.Resources.Energy.Quantity) * .5)
			}
		},
		Resolve: func() {},
		Expire: func() {
			if planet, ok := targetEntity.(*models.Planet); ok {
				planet.Resources.Minerals.Quantity -= int(float64(planet.Resources.Minerals.Quantity) * .20)
				planet.Resources.Food.Quantity -= int(float64(planet.Resources.Food.Quantity) * .20)
				planet.Resources.Energy.Quantity -= int(float64(planet.Resources.Energy.Quantity) * .20)
			}
		},
	}
}
