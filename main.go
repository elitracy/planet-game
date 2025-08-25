package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/elitracy/planets/components"
	. "github.com/elitracy/planets/planet"
	. "github.com/elitracy/planets/system"
	ui "github.com/gizak/termui/v3"
)

const NUM_SYSTEMS = 3
const MIN_PLANETS = 2
const MAX_PLANETS = 5

func main() {

	// init game
	// generate systems
	var systems []System

	for range NUM_SYSTEMS {
		system := GenerateSystem()

		for range rand.Intn(MAX_PLANETS-MIN_PLANETS) + MIN_PLANETS {
			system.Planets = append(system.Planets, GeneratePlanet())
		}

		systems = append(systems, system)
	}

	for _, s := range systems {
		fmt.Println(s)
	}

	// main game loop
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	systemTree := &components.SystemTree{Systems: systems}
	systemTree.Init()
	systemTree.Update()

	planetPopBC := &components.PopulationBarChart{System: systems[0]}
	planetPopBC.Init()
	planetPopBC.Update()

	planetBPopBC := &components.PopulationBarChart{System: systems[1]}
	planetBPopBC.Init()
	planetBPopBC.Update()

	dashboard := &Dashboard{
		Components: [][]Widget{
			{systemTree, planetBPopBC},
			{planetPopBC},
		},
	}
	dashboard.SetRects()
	dashboard.SetActiveWidget(0, 0)
	dashboard.Render()

	uiEvents := ui.PollEvents()

	for {
		dashboard.Render()
		e := <-uiEvents
		switch e.ID {
		case "q", "C-c":
			return
		default:
			dashboard.HandleKey(e.ID)
		}
	}

}
