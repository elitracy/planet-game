package models

import (
	core "github.com/elitracy/planets/core"
)

type StarSystem struct {
	*CoreEntity
	Planets   []*Planet
	Scouted   bool
	Colonized bool
}

func CreateStarSystem(name string, planets []*Planet, position core.Position) *StarSystem {
	return &StarSystem{
		CoreEntity: &CoreEntity{
			Name:     name,
			Position: position,
		},
		Planets:   planets,
		Colonized: false,
	}
}

func (s *StarSystem) Tick() {
	for _, planet := range s.Planets {
		planet.Tick()
		if planet.Colonized {
			s.Colonized = true
		}
	}
}
