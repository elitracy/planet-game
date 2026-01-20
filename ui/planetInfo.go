package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
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
	case paneResizeMsg:
		logging.Info("planet info resize msg: %v", msg)
		p.height = msg.height
		p.width = msg.width
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
			return p, pushDetailStackCmd(childPaneID)
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
			return p, pushDetailStackCmd(childPaneID)
		case "esc":
			PaneManager.RemovePane(childPaneID)
			return p, popDetailStackCmd(p.Pane.id)
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

	resources := "Resources:\n"
	resources += fmt.Sprintf("Food:     %d\n", p.planet.Resources.Food.GetQuantity())
	resources += fmt.Sprintf("Minerals: %d\n", p.planet.Resources.Minerals.GetQuantity())
	resources += fmt.Sprintf("Energy:   %d", p.planet.Resources.Energy.GetQuantity())

	constructions := "Constructions:\n"
	constructions += fmt.Sprintf("Farms:       %d\n", len(p.planet.Constructions.Farms))
	constructions += fmt.Sprintf("Mines:       %d\n", len(p.planet.Constructions.Mines))
	constructions += fmt.Sprintf("Solar Grids: %d", len(p.planet.Constructions.SolarGrids))

	title = consts.Style.Width(p.width).Align(lipgloss.Center).Bold(true).Render(title)

	info := lipgloss.JoinHorizontal(lipgloss.Top, resources, constructions)
	info = consts.Style.
		Render(info)

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, title, population, info)

	colonizeButton := "Order Colonization: O"
	scoutButton := "Scout: S"
	changeAllocationsButton := "Change Allocations: A"

	buttons := lipgloss.JoinHorizontal(lipgloss.Left, colonizeButton, " | ", scoutButton, " | ", changeAllocationsButton)
	buttons = consts.Style.Width(p.width).Border(lipgloss.NormalBorder(), true, false, false, false).Render(buttons)

	logging.Info("height: %v", p.height)
	logging.Info("width: %v", p.width)
	topContent := lipgloss.Place(
		p.width,
		p.height-lipgloss.Height(buttons),
		lipgloss.Left,
		lipgloss.Top,
		infoContainer,
	)
	content := lipgloss.JoinVertical(lipgloss.Left, topContent, buttons)

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
