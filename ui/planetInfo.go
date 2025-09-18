package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
)

type PlanetInfoPane struct {
	Pane
	id     int
	title  string
	planet *models.Planet
}

func (p *PlanetInfoPane) Init() tea.Cmd { return nil }
func (p *PlanetInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "c":
			pane := NewCreateColonyPane(
				"Colonize: "+p.planet.Name,
				1000,
				p.planet,
			)
			PushFocus(pane)
			return ActivePane(), nil

		case "esc":
			return PopFocus(), nil
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	default:
	}
	return p, nil
}

func (p *PlanetInfoPane) View() string {
	title := p.planet.Name
	if p.planet.ColonyName != "" {
		title += Theme.blurredStyle.Render(fmt.Sprintf(" [%s]", p.planet.ColonyName))
	}

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
		PaddingRight(1).
		Inherit(defaultStyle).
		Render(resources)
	constructions = defaultStyle.Render(constructions)

	info := lipgloss.JoinHorizontal(lipgloss.Top, resources, constructions)
	info = lipgloss.NewStyle().
		Render(info)

	infoContainer := lipgloss.JoinVertical(lipgloss.Center, title, population, info)

	colonizeButton := Theme.focusedStyle.Underline(true).Render("C") + "olonize"
	colonizeButton = lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(colonizeButton)

	changeAllocationsButton := "Change " + Theme.focusedStyle.Underline(true).Render("A") + "llocations"
	changeAllocationsButton = lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(changeAllocationsButton)

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, colonizeButton, changeAllocationsButton)

	content := lipgloss.JoinVertical(lipgloss.Center, infoContainer, buttons)

	return content

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
