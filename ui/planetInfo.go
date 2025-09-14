package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
)

type PlanetInfoPane struct {
	BasePane
	id     int
	title  string
	planet *models.Planet
}

func (p *PlanetInfoPane) Init() tea.Cmd { return tick(p.id) }
func (p *PlanetInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		if msg.id == p.id {
			return p, tick(p.id)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return PopFocus(), nil
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *PlanetInfoPane) View() string {
	title := fmt.Sprintf("%s", p.planet.Name)

	population := fmt.Sprintf("Population: %d", p.planet.Population)

	// Resources
	resources := "Resources:\n"
	resources += fmt.Sprintf("Food:     %d\n", p.planet.Resources.Food.GetQuantity())
	resources += fmt.Sprintf("Minerals: %d\n", p.planet.Resources.Minerals.GetQuantity())
	resources += fmt.Sprintf("Energy:   %d", p.planet.Resources.Energy.GetQuantity())

	// Constructions
	constructions := "Constructions:\n"
	constructions += fmt.Sprintf("Farms:       %d\n", len(p.planet.Constructions.Farms))
	constructions += fmt.Sprintf("Mines:       %d\n", len(p.planet.Constructions.Mines))
	constructions += fmt.Sprintf("Solar Grids: %d", len(p.planet.Constructions.SolarGrids))

	defaultStyle := lipgloss.NewStyle().
		Padding(1).
		PaddingTop(0)

	title = lipgloss.NewStyle().Bold(true).Render(title)
	title = defaultStyle.Render(title)

	population = defaultStyle.Render(population)

	resources = lipgloss.NewStyle().
		BorderRight(true).
		Inherit(defaultStyle).
		Render(resources)
	constructions = defaultStyle.Render(constructions)

	info := lipgloss.JoinHorizontal(lipgloss.Top, resources, constructions)
	info = lipgloss.NewStyle().
		PaddingTop(1).
		Render(info)

	container := lipgloss.JoinVertical(lipgloss.Center, title, population, info)

	return container

}

func (p PlanetInfoPane) GetId() int       { return p.id }
func (p PlanetInfoPane) GetTitle() string { return p.title }

func NewPlanetInfoPane(title string, id int, planet *models.Planet) *PlanetInfoPane {

	return &PlanetInfoPane{
		title:  title,
		id:     id,
		planet: planet,
	}
}
