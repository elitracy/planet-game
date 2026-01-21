package models

import (
	"fmt"

	. "github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/interfaces"
	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/resources"
	"github.com/elitracy/planets/models/stabilities"
)

type Planet struct {
	ID                   int
	Name                 string
	ColonyName           string
	Population           int
	PopulationGrowthRate int
	Resources
	Stabilities
	Constructions
	Position
	OrderQueue []Event
}

func (p Planet) GetID() int            { return p.ID }
func (p Planet) GetName() string       { return p.Name }
func (p Planet) GetPosition() Position { return p.Position }
func (p Planet) GetOrders() []Event    { return p.OrderQueue }

func (p *Planet) PushOrder(order Event) {
	p.OrderQueue = append(p.OrderQueue, order)
}

func (p *Planet) PopOrder() Event {
	if len(p.OrderQueue) == 0 {
		return nil
	}

	order := p.OrderQueue[0]
	p.OrderQueue = p.OrderQueue[1:]
	return order
}

func (p Planet) GetTotalFarmProduction() int {
	total_rate := 0

	for _, p := range p.Constructions.Farms {
		total_rate += p.GetProductionRate()
	}

	return total_rate
}

func (p Planet) GetTotalMineProduction() int {
	total_rate := 0
	for _, p := range p.Constructions.Mines {
		total_rate += p.GetProductionRate()
	}

	return total_rate
}

func (p Planet) GetTotalSolarGridProduction() int {
	total_rate := 0

	for _, p := range p.Constructions.SolarGrids {
		total_rate += p.GetProductionRate()
	}

	return total_rate
}

func CreatePlanet(name string, x, y, z, pop, initial_food, initial_mineral, intital_energy, initial_food_rate, initial_mineral_rate, initial_energy_rate, num_farms, num_mines, num_solar_grids int) Planet {
	planet := Planet{
		Name:       name,
		Population: pop,
		Position:   Position{X: x, Y: y, Z: z},
		Resources: Resources{
			Food: resources.Food{
				Quantity:        initial_food,
				ConsumptionRate: initial_food_rate,
			},
			Minerals: resources.Mineral{
				Quantity:        initial_mineral,
				ConsumptionRate: initial_mineral_rate,
			},
			Energy: resources.Energy{
				Quantity:        intital_energy,
				ConsumptionRate: initial_energy_rate,
			},
		},
	}

	for range num_farms {
		planet.Constructions.Farms = append(planet.Constructions.Farms, constructions.CreateFarm(1))
	}

	for range num_mines {
		planet.Constructions.Mines = append(planet.Constructions.Mines, constructions.CreateMine(1))
	}

	for range num_solar_grids {
		planet.Constructions.SolarGrids = append(planet.Constructions.SolarGrids, constructions.CreateSolarGrid(1))
	}

	return planet
}

type Resources struct {
	Food     resources.Food
	Minerals resources.Mineral
	Energy   resources.Energy
}

func (r Resources) String() string {
	var output string

	output += fmt.Sprintf("| Food:       %d @ %d\n", r.Food.GetQuantity(), r.Food.GetConsumptionRate())
	output += fmt.Sprintf("| Energy:     %d @ %d\n", r.Energy.GetQuantity(), r.Energy.GetConsumptionRate())
	output += fmt.Sprintf("| Minerals:   %d @ %d", r.Minerals.GetQuantity(), r.Minerals.GetConsumptionRate())

	return output
}

type Stabilities struct {
	Corruption stabilities.Corruption
	Happiness  stabilities.Happiness
	Unrest     stabilities.Unrest
}

func (s Stabilities) String() string {
	var output string

	output += fmt.Sprintf("| Corruption: %.2f @ %.2f\n", s.Corruption.Quantity, s.Corruption.GetGrowthRate())
	output += fmt.Sprintf("| Happiness:  %.2f @ %.2f\n", s.Happiness.Quantity, s.Happiness.GetGrowthRate())
	output += fmt.Sprintf("| Unrest:     %.2f @ %.2f", s.Unrest.Quantity, s.Unrest.GetGrowthRate())

	return output
}

type Constructions struct {
	Farms      []constructions.Farm
	Mines      []constructions.Mine
	SolarGrids []constructions.SolarGrid
}

func (c Constructions) String() string {
	var output string

	output += "| Farms"
	for i, f := range c.Farms {
		output += fmt.Sprintf("\n 🌾 Farm[%d]: %d @ %d", i, f.GetQuantity(), f.GetProductionRate())
	}

	output += "| Mines\n"
	for i, m := range c.Mines {
		output += fmt.Sprintf("\n ⛏️ Mine[%d]: %d @ %d", i, m.GetQuantity(), m.GetProductionRate())
	}

	output += "| Solar Grids\n"
	for i, sg := range c.SolarGrids {
		output += fmt.Sprintf("\n ☀️ Solar Grid[%d]: %d @ %d", i, sg.GetQuantity(), sg.GetProductionRate())
	}

	return output
}
