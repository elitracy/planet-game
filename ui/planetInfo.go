package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/state"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/orders"
)

type PlanetInfoPane struct {
	*Pane

	childPaneID core.PaneID
	planet      *models.Planet
}

func (p *PlanetInfoPane) Init() tea.Cmd { return nil }
func (p *PlanetInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var childPaneID core.PaneID

	switch msg := msg.(type) {
	case core.TickMsg:
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
				&state.State.ShipManager,
				func(ship *models.Ship) {
					order := orders.NewScoutShipOrder(ship, p.planet.Position, state.State.Tick+40)
					state.State.OrderScheduler.Push(order)
				},
			)

			childPaneID = PaneManager.AddPane(pane)
			return p, pushFocusCmd(childPaneID)
		case "esc":
			PaneManager.RemovePane(childPaneID)
			return p, popFocusCmd(p.Pane.id)
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
		title += consts.Theme.BlurredStyle.Render(fmt.Sprintf(" [%s]", p.planet.ColonyName))
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

	defaultStyle := consts.Style.
		Padding(1).
		PaddingTop(0)

	title = consts.Style.Bold(true).Render(title)
	title = defaultStyle.Render(title)

	population = defaultStyle.Render(population)

	resources = consts.Style.
		PaddingRight(1).
		Inherit(defaultStyle).
		Render(resources)
	constructions = defaultStyle.Render(constructions)

	info := lipgloss.JoinHorizontal(lipgloss.Top, resources, constructions)
	info = consts.Style.
		Render(info)

	infoContainer := lipgloss.JoinVertical(lipgloss.Center, title, population, info)

	colonizeButton := consts.Theme.FocusedStyle.Underline(true).Render("O") + "rder Colonization"
	colonizeButton = consts.Style.
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(colonizeButton)

	scoutButton := consts.Theme.FocusedStyle.Underline(true).Render("S") + "cout"
	scoutButton = consts.Style.
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(scoutButton)

	changeAllocationsButton := "Change " + consts.Theme.FocusedStyle.Underline(true).Render("A") + "llocations"
	changeAllocationsButton = consts.Style.
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render(changeAllocationsButton)

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, colonizeButton, scoutButton, changeAllocationsButton)

	content := lipgloss.JoinVertical(lipgloss.Center, infoContainer, buttons)

	return content

}

func NewPlanetInfoPane(title string, planet *models.Planet) *PlanetInfoPane {

	return &PlanetInfoPane{
		Pane: &Pane{
			title: title,
		},
		planet: planet,
	}
}
