package models

import (
	"fmt"

	. "github.com/elitracy/planets/core"
)

type StarSystem struct {
	Name    string
	Planets []*Planet
	Position
}

func (s *StarSystem) String() string {
	var output string

	output += fmt.Sprintf("%s %v\n", s.Name, s.Position)
	for _, p := range s.Planets {

		output += "--------------------------------------------------\n"
		output += fmt.Sprintf("%s %v\n", p.Name, p.Position)
		output += fmt.Sprintf("| Resources:\n")
		output += fmt.Sprintf("|-| > Food:    %d\n", p.Food.Quantity)
		output += fmt.Sprintf("  | > Mineral: %d\n", p.Minerals.Quantity)
		output += fmt.Sprintf("  | > Energy:  %d\n", p.Energy.Quantity)
		output += fmt.Sprintf("| Constructions:\n")
		output += fmt.Sprintf("|-| > Farms:        %d @ %d/t\n", len(p.Farms), p.GetTotalFarmProduction())
		output += fmt.Sprintf("  | > Mines:        %d @ %d/t\n", len(p.Mines), p.GetTotalMineProduction())
		output += fmt.Sprintf("  | > Solar Grids:  %d @ %d/t\n", len(p.SolarGrids), p.GetTotalSolarGridProduction())
		output += fmt.Sprintf("| Stabilities:\n")
		output += fmt.Sprintf("|-| > Unrest:      %.2f @ %.2f/t\n", p.Unrest.GetQuantity(), p.Unrest.GetGrowthRate())
		output += fmt.Sprintf("  | > Happiness:   %.2f @ %.2f/t\n", p.Happiness.GetQuantity(), p.Happiness.GetGrowthRate())
		output += fmt.Sprintf("  | > Corruption:  %.2f @ %.2f/t\n", p.Corruption.GetQuantity(), p.Corruption.GetGrowthRate())
		output += "--------------------------------------------------\n"
	}

	return output
}
