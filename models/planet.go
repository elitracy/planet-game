package models

import (
	"fmt"

	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/resources"
	"github.com/elitracy/planets/models/stabilities"
)

type Planet struct {
	Name                 string
	Popluation           int
	PopulationGrowthRate int
	Resources            PlanetResources
	Stabilities          PlanetStabilities
	Constructions        PlanetConstructions
	Location             Location
}

func (p Planet) String() string {
	var output string

	output += fmt.Sprintf("🪐 %s [%d, %d]\n", p.Name, p.Location.Coordinates.X, p.Location.Coordinates.Y)
	output += fmt.Sprintf("|  Food:     %d @ %d\n", p.Resources.Food.Quantity, p.Resources.Food.ConsumptionRate)
	output += fmt.Sprintf("|  Energy:   %d @ %d\n", p.Resources.Energy.Quantity, p.Resources.Energy.ConsumptionRate)
	output += fmt.Sprintf("|  Minerals: %d @ %d", p.Resources.Minerals.Quantity, p.Resources.Minerals.ConsumptionRate)

	return output
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
	Farm      []constructions.Farm
	Mine      []constructions.Mine
	SolarGrid []constructions.SolarGrid
}
