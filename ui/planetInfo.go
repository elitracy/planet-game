package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/orders"
	. "github.com/elitracy/planets/state"
)

type PlanetInfoPane struct {
	Pane
	id          int
	childPaneID int
	title       string
	planet      *Planet
}

func (p *PlanetInfoPane) Init() tea.Cmd { return nil }
func (p *PlanetInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var childPaneID int

	switch msg := msg.(type) {
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "o":
			pane := NewCreateColonyPane(
				"Order Colonization: "+p.planet.Name,
				p.planet,
			)

			childPaneID = PaneManager.AddPane(pane)
			return p, pushFocusCmd(childPaneID)
		case "s":
			pane := CreateNewShipManagementPane(
				"Ship Management",
				&State.ShipManager,
				func(ship *Ship) {
					order := orders.NewScoutShipOrder(ship, p.planet.Position, State.Tick+40)
					State.OrderScheduler.Push(order)
				},
			)

			childPaneID = PaneManager.AddPane(pane)
			return p, pushFocusCmd(childPaneID)
		case "esc":
			PaneManager.RemovePane(childPaneID)
			return p, popFocusCmd()
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

	defaultStyle := Style.
		Padding(1).
		PaddingTop(0)

	title = Style.Bold(true).Render(title)
	title = defaultStyle.Render(title)

	population = defaultStyle.Render(population)

	resources = Style.
		PaddingRight(1).
		Inherit(defaultStyle).
		Render(resources)
	constructions = defaultStyle.Render(constructions)

	info := lipgloss.JoinHorizontal(lipgloss.Top, resources, constructions)
	info = Style.
		Render(info)

	infoContainer := lipgloss.JoinVertical(lipgloss.Center, title, population, info)

	colonizeButton := Theme.focusedStyle.Underline(true).Render("O") + "rder Colonization"
	colonizeButton = Style.
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(colonizeButton)

	scoutButton := Theme.focusedStyle.Underline(true).Render("S") + "cout"
	scoutButton = Style.
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(scoutButton)

	changeAllocationsButton := "Change " + Theme.focusedStyle.Underline(true).Render("A") + "llocations"
	changeAllocationsButton = Style.
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(changeAllocationsButton)

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, colonizeButton, scoutButton, changeAllocationsButton)

	content := lipgloss.JoinVertical(lipgloss.Center, infoContainer, buttons)

	return content

}

func (p PlanetInfoPane) GetId() int       { return p.id }
func (p *PlanetInfoPane) SetId(id int)    { p.id = id }
func (p PlanetInfoPane) GetTitle() string { return p.title }

func NewPlanetInfoPane(title string, planet *Planet) *PlanetInfoPane {

	return &PlanetInfoPane{
		title:  title,
		planet: planet,
	}
}
