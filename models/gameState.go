package models

import (
	"github.com/elitracy/planets/models/resources"
)

type GameState struct {
	CurrentTick int
	StarSystems []StarSystem
	MessagePayLoads []Payload[string]
	ResourcePayLoads []Payload[resources.Resource]
}
