package models

import (
	core "github.com/elitracy/planets/core"
)

type StarSystem struct {
	Name    string
	Planets []*Planet
	Colonized bool
	core.Position
}

func CreateStarSystem(name string, planets []*Planet, position core.Position) *StarSystem {
	return &StarSystem{
		Name: name,
		Planets: planets,
		Colonized: false,
		Position: position,
	}
}
