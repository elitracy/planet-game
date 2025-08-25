package models

import (
	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/resources"
	"github.com/elitracy/planets/models/stabilities"
)

type Planet struct {
	Name          string
	Popluation    int
	PopulationGrowthRate    int
	Resources     PlanetResources
	Stabilities   PlanetStabilities
	Constructions PlanetConstructions
	Location      Coordinates
}

type PlanetResources struct {
	Food     resources.Food
	Minerals resources.Mineral
	Energy   resources.Energy
}

type PlanetStabilities struct {
	Corruption stabilities.Corruption
	Happiness  stabilities.Happiness
	Unrest     stabilities.Unrest
}

type PlanetConstructions struct {
	Farm      constructions.Farm
	Mine      constructions.Mine
	SolarGrid constructions.SolarGrid
}
