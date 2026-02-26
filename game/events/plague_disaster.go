package events

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/models"
)

func NewPlagueDisaster(targetEntity models.Entity, tick engine.Tick) *Event {
	return &Event{
		Name:        "Plague Disaster",
		Description: fmt.Sprintf("A plague has fallen upon %v", targetEntity.GetName()),
		Severity:    Moderate,
		Target:      targetEntity,
		Resolved:    false,
		Effect: func() {
			if planet, ok := targetEntity.(*models.Planet); ok {
				planet.Population -= int(float64(planet.Population) * 0.01)
				planet.PopulationGrowthRate -= 100
			}
		},
		Expire:   func() {},
		Duration: config.TICKS_PER_PULSE * 30,
		Interval: config.TICKS_PER_PULSE,
	}
}
