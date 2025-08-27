package models

import (
	"fmt"

	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/resources"
	"github.com/elitracy/planets/models/stabilities"
)

type Planet struct {
	Name                 string
	Population           int
	PopulationGrowthRate int
	Resources            PlanetResources
	Stabilities          PlanetStabilities
	Constructions        PlanetConstructions
	Location             Location
}

func (p Planet) String() string {
	var output string

	output += fmt.Sprintf("🪐 %s [%d, %d]\n", p.Name, p.Location.Coordinates.X, p.Location.Coordinates.Y)
	output += fmt.Sprintf("| Population: %d @ %d\n", p.Population, p.PopulationGrowthRate)
	output += fmt.Sprintf(p.Resources.String() + "\n")
	output += fmt.Sprintf(p.Constructions.String() + "\n")
	output += fmt.Sprintf(p.Stabilities.String() + "\n")

	return output
}

type PlanetResources struct {
	Food     resources.Food
	Minerals resources.Mineral
	Energy   resources.Energy
}

func (p PlanetResources) String() string {
	var output string

	output += fmt.Sprintf("| Food:       %d @ %d\n", p.Food.GetQuantity(), p.Food.GetConsumptionRate())
	output += fmt.Sprintf("| Energy:     %d @ %d\n", p.Energy.GetQuantity(), p.Energy.GetConsumptionRate())
	output += fmt.Sprintf("| Minerals:   %d @ %d", p.Minerals.GetQuantity(), p.Minerals.GetConsumptionRate())

	return output
}

type PlanetStabilities struct {
	Corruption stabilities.Corruption
	Happiness  stabilities.Happiness
	Unrest     stabilities.Unrest
}

func (p PlanetStabilities) String() string {
	var output string

	output += fmt.Sprintf("| Corruption: %.2f @ %.2f\n", p.Corruption.Quantity, p.Corruption.GetGrowthRate())
	output += fmt.Sprintf("| Happiness:  %.2f @ %.2f\n", p.Happiness.Quantity, p.Happiness.GetGrowthRate())
	output += fmt.Sprintf("| Unrest:     %.2f @ %.2f", p.Unrest.Quantity, p.Unrest.GetGrowthRate())

	return output
}

type PlanetConstructions struct {
	Farms      []constructions.Farm
	Mines      []constructions.Mine
	SolarGrids []constructions.SolarGrid
}

func (p PlanetConstructions) String() string {
	var output string

	output += fmt.Sprintf("| Farms\n")
	for i, f := range p.Farms {
		output += fmt.Sprintf(" 🌾 Farm[%d]: %d @ %d\n", i, f.GetQuantity(), f.GetProductionRate())
	}

	output += fmt.Sprintf("| Mines\n")
	for i, m := range p.Mines {
		output += fmt.Sprintf(" ⛏️ Mine[%d]: %d @ %d\n", i, m.GetQuantity(), m.GetProductionRate())
	}

	output += fmt.Sprintf("| Solar Grids\n")
	for i, sg := range p.SolarGrids {
		output += fmt.Sprintf(" ☀️ Solar Grid[%d]: %d @ %d\n", i, sg.GetQuantity(), sg.GetProductionRate())
	}

	return output
}
