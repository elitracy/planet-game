package events

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/models"
)

func NewMineralDiscovery(targetEntity models.Entity, tick engine.Tick) *Event {
	return &Event{
		Name:        "Mineral Discovery",
		Description: fmt.Sprintf("A mineral deposit has been found on an asteroid!"),
		Severity:    Moderate,
		Target:      targetEntity,
		Resolved:    false,
		Effect: func() {
			if planet, ok := targetEntity.(*models.Planet); ok {
				planet.Resources.Minerals.Quantity += 50_000
			}
		},
		Expire:   func() {},
		Duration: 0,
		Interval: 0,
	}
}
