package main

import (
	"github.com/elitracy/planets/engine"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/resources"
)

func main() {

	gameState := GameState{}

	systemA := &StarSystem{}
	systemB := &StarSystem{}
	earth := &Planet{
		Name:                 "Earth",
		Popluation:           1000,
		PopulationGrowthRate: 5,
		Resources: PlanetResources{
			Food: resources.Food{
				Quantity:        5000,
				ConsumptionRate: 1,
			},
			Minerals: resources.Mineral{
				Quantity:        5000,
				ConsumptionRate: 1,
			},
			Energy: resources.Energy{
				Quantity:        5000,
				ConsumptionRate: 1,
			},
		},
	}

	james := &Planet{
		Name: "James",
		Location: Location{
			Coordinates: Coordinates{X: 100, Y: 500},
		},
		Popluation:           1000,
		PopulationGrowthRate: 5,
		Resources: PlanetResources{
			Food: resources.Food{
				Quantity:        5000,
				ConsumptionRate: 1,
			},
			Minerals: resources.Mineral{
				Quantity:        5000,
				ConsumptionRate: 1,
			},
			Energy: resources.Energy{
				Quantity:        5000,
				ConsumptionRate: 1,
			},
		},
	}

	systemA.Planets = append(systemA.Planets, earth)
	systemB.Planets = append(systemB.Planets, james)

	gameState.StarSystems = append(gameState.StarSystems, systemA, systemB)

	engine.RunGame(&gameState)
}
