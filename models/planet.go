package models

import (
	. "github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/interfaces"
	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/resources"
	"github.com/elitracy/planets/models/stabilities"
)

const (
	STARTING_FOOD                     = 1000
	STARTING_FOOD_CONSUMPTION_RATE    = 1
	STARTING_MINERAL                  = 1000
	STARTING_MINERAL_CONSUMPTION_RATE = 1
	STARTING_ENERGY                   = 1000
	STARTING_ENERGY_CONSUMPTION_RATE  = 1
)

type Planet struct {
	ID                   int
	Name                 string
	Population           int
	PopulationGrowthRate int
	Colonized            bool

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

func (p Planet) GetFarmProduction() int {
	total_rate := 0
	for _, p := range p.Constructions.Farms {
		total_rate += p.GetProductionRate()
	}

	return total_rate
}

func (p Planet) GetMineProduction() int {
	total_rate := 0
	for _, p := range p.Constructions.Mines {
		total_rate += p.GetProductionRate()
	}

	return total_rate
}

func (p Planet) GetSolarGridProduction() int {
	total_rate := 0
	for _, p := range p.Constructions.SolarGrids {
		total_rate += p.GetProductionRate()
	}

	return total_rate
}

func CreatePlanet(name string, x, y, z, pop, num_farms, num_mines, num_solar_grids int) Planet {
	planet := Planet{
		Name:       name,
		Population: pop,
		Position:   Position{X: x, Y: y, Z: z},
		Resources: Resources{
			Food: resources.Food{
				Quantity:        STARTING_FOOD,
				ConsumptionRate: STARTING_FOOD_CONSUMPTION_RATE,
			},
			Minerals: resources.Mineral{
				Quantity:        STARTING_MINERAL,
				ConsumptionRate: STARTING_MINERAL_CONSUMPTION_RATE,
			},
			Energy: resources.Energy{
				Quantity:        STARTING_ENERGY,
				ConsumptionRate: STARTING_ENERGY_CONSUMPTION_RATE,
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

type Stabilities struct {
	Corruption stabilities.Corruption
	Happiness  stabilities.Happiness
	Unrest     stabilities.Unrest
}

type Constructions struct {
	Farms      []constructions.Farm
	Mines      []constructions.Mine
	SolarGrids []constructions.SolarGrid
}
