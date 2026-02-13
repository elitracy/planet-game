package models

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/models/constructions"
	"github.com/elitracy/planets/game/models/resources"
	"github.com/elitracy/planets/game/models/stabilities"
)

const (
	STARTING_FOOD                     = 1000
	STARTING_FOOD_CONSUMPTION_RATE    = 1
	STARTING_MINERAL                  = 10000
	STARTING_MINERAL_CONSUMPTION_RATE = 1
	STARTING_ENERGY                   = 10000
	STARTING_ENERGY_CONSUMPTION_RATE  = 1
)

type Planet struct {
	*CoreEntity
	Population           int
	PopulationGrowthRate int
	Colonized            bool
	Scouted              bool

	Resources
	Stabilities
	Constructions
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

func CreatePlanet(name string, position engine.Position, pop, num_farms, num_mines, num_solar_grids int) Planet {
	planet := Planet{
		CoreEntity: &CoreEntity{
			Name:     name,
			Location: Location{Position: position},
		},
		Population:           pop,
		PopulationGrowthRate: config.POPULATION_GROWTH_PER_PULSE,
		Resources: Resources{
			Food: resources.Food{
				Quantity:        pop * config.FOOD_PER_PERSON * config.NUM_DAYS_FED,
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
		Stabilities: Stabilities{
			Happiness: stabilities.Happiness{
				Quantity:   1,
				GrowthRate: 0,
			},
			Corruption: stabilities.Corruption{
				Quantity:   0,
				GrowthRate: 0,
			},
			Unrest: stabilities.Unrest{
				Quantity:   0,
				GrowthRate: 0,
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

func (p *Planet) PushOrder(order task.Task) {
	p.OrderQueue = append(p.OrderQueue, order)
}

func (p *Planet) PopOrder() task.Task {
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

func (p *Planet) Tick() {
	p.TickResources()
	p.TickStabilities()
	p.TickConstructions()

}

func (p *Planet) TickResources() {

	currentFood := p.Food.Quantity
	requiredFood := p.Population * config.FOOD_PER_PERSON

	if currentFood >= requiredFood {
		p.Food.Quantity -= requiredFood

		currentFood = p.Food.Quantity
		addPopRequiredFood := config.FOOD_PER_PERSON * p.PopulationGrowthRate

		if currentFood >= addPopRequiredFood {
			p.Food.Quantity -= config.FOOD_PER_PERSON * p.PopulationGrowthRate
			p.Population += p.PopulationGrowthRate
		}
	}
}

func (p *Planet) TickStabilities() {

	currentFood := p.Food.Quantity
	requiredFood := p.Population * config.FOOD_PER_PERSON

	if currentFood < requiredFood {
		p.Stabilities.Happiness.GrowthRate -= .05
	}

	p.Happiness.Tick()
	p.Corruption.Tick()
	p.Unrest.Tick()
}

func (p *Planet) TickConstructions() {

	for _, farm := range p.Constructions.Farms {
		p.Resources.Food.Quantity += farm.Quantity * farm.ProductionRate
	}

	for _, mine := range p.Constructions.Mines {
		p.Resources.Minerals.Quantity += mine.Quantity * mine.ProductionRate
	}

	for _, solarGrid := range p.Constructions.SolarGrids {
		p.Resources.Energy.Quantity += solarGrid.Quantity * solarGrid.ProductionRate
	}
}
