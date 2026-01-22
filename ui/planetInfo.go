package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/orders"
)

type PlanetInfoPane struct {
	*Pane

	childPaneID core.PaneID
	planet      *models.Planet
	theme       UITheme
}

func (p *PlanetInfoPane) Init() tea.Cmd { return nil }
func (p *PlanetInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var paneID core.PaneID

	switch msg := msg.(type) {
	case paneResizeMsg:
		p.height = msg.height
		p.width = msg.width
	case tea.KeyMsg:
		switch msg.String() {
		case "o":
			pane := NewCreateColonyPane(
				"Order Colonization: "+p.planet.Name,
				p.planet,
			)

			paneID = PaneManager.AddPane(pane)
			return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
		case "s":
			pane := CreateNewShipManagementPane(
				"Ship Management",
				&state.State.ShipManager,
				func(ship *models.Ship) {
					order := orders.NewScoutShipOrder(ship, p.planet.Position, state.State.Tick+40)
					state.State.OrderScheduler.Push(order)
				},
			)

			paneID = PaneManager.AddPane(pane)
			return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
		case "esc":
			PaneManager.RemovePane(paneID)
			return p, tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *PlanetInfoPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.planet.Name

	population := fmt.Sprintf("Population: %d", p.planet.Population)

	resources := fmt.Sprintf("\nFood: %10d", p.planet.Food.GetQuantity())
	resources += fmt.Sprintf("\nMinerals: %10d", p.planet.Minerals.GetQuantity())
	resources += fmt.Sprintf("\nEnergy:  %10d", p.planet.Energy.GetQuantity())

	constructions := fmt.Sprintf("\nFarms: %10d", len(p.planet.Farms))
	constructions += fmt.Sprintf("\nMines: %10d", len(p.planet.Mines))
	constructions += fmt.Sprintf("\nSolar Grids: %10d", len(p.planet.SolarGrids))

	stabilities := fmt.Sprintf("\nHappiness: %10.0f%%", p.planet.Happiness.GetQuantity()*100)
	stabilities += fmt.Sprintf("\nCorruption: %10.0f%%", p.planet.Corruption.GetQuantity()*100)
	stabilities += fmt.Sprintf("\nUnrest: %10.0f%%", p.planet.Unrest.GetQuantity()*100)

	title = Style.Width(p.width).Align(lipgloss.Center).Bold(true).Render(title)

	info := lipgloss.JoinVertical(lipgloss.Left, resources, constructions, stabilities)
	info = Style.Render(info)

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, title, population, info)

	colonizeButton := "Order Colonization: O"
	scoutButton := "Scout: S"
	changeAllocationsButton := "Change Allocations: A"

	buttons := lipgloss.JoinHorizontal(lipgloss.Left, colonizeButton, " | ", scoutButton, " | ", changeAllocationsButton)
	buttons = Style.Width(p.width).Border(lipgloss.NormalBorder(), true, false, false, false).Render(buttons)

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
