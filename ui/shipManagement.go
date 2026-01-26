package ui

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/state"
)

type ShipManagementPane struct {
	*Pane
	cursor        int
	currentShipID int
	sortedShips   []*models.Ship
	manager       *models.ShipManager
	OnSelect      func(ship *models.Ship)
	theme         UITheme
}

func CreateNewShipManagementPane(title string, shipManager *models.ShipManager, callback func(ship *models.Ship)) *ShipManagementPane {
	pane := &ShipManagementPane{
		Pane: &Pane{
			title: title,
		},
		manager:  shipManager,
		OnSelect: callback,
	}

	return pane
}

func (p *ShipManagementPane) Init() tea.Cmd {
	for _, ship := range p.manager.Ships {
		p.sortedShips = append(p.sortedShips, ship)
	}
	sort.Slice(p.sortedShips, func(i, j int) bool {
		return p.sortedShips[i].Name > p.sortedShips[j].Name
	})

	return nil
}

func (p *ShipManagementPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.Pane.id {
			p.Pane.width = msg.width - 2
			p.Pane.height = msg.height

			return p, nil
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			p.OnSelect(state.State.ShipManager.Ships[p.currentShipID])
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.sortedShips)-1 {
				p.cursor++
			}
		case "esc":
			return p, tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *ShipManagementPane) View() string {
	p.theme = GetPaneTheme(p)

	var rows []string

	idx := 0
	for _, ship := range p.sortedShips {
		var row string
		if idx == p.cursor {
			p.currentShipID = ship.GetID()

			row = fmt.Sprintf("%s - %s", ship.GetName(), ship.GetPosition())
			row = Style.Border(lipgloss.NormalBorder()).Padding(0, 1).Render(row)
			row = p.theme.FocusedStyle.Render(row)
		} else {
			row = fmt.Sprintf("%s - %s", ship.GetName(), ship.GetPosition())
			row = Style.Border(lipgloss.NormalBorder()).Padding(0, 1).Render(row)
			row = p.theme.BlurredStyle.Render(row)
		}

		rows = append(rows, row)
		idx++
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	title := "Ships"
	title = Style.Width(p.Pane.width).AlignHorizontal(lipgloss.Center).Render(title)

	content = lipgloss.JoinVertical(lipgloss.Left, title, content)

	return content

}
