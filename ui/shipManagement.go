package ui

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	. "github.com/elitracy/planets/models"
)

type ShipManagementPane struct {
	Pane
	id            int
	title         string
	width         int
	height        int
	cursor        int
	currentShipID int
	sortedShips   []*Ship
	manager       *ShipManager
	OnSelect      func(ship *Ship)
}

func CreateNewShipManagementPane(title string, shipManager *ShipManager, callback func(ship *Ship)) *ShipManagementPane {
	pane := &ShipManagementPane{
		title:    title,
		manager:  shipManager,
		OnSelect: callback,
	}

	return pane
}

func (p ShipManagementPane) GetId() int       { return p.id }
func (p *ShipManagementPane) SetId(id int)    { p.id = id }
func (p ShipManagementPane) GetTitle() string { return p.title }

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
		if msg.paneID == p.GetId() {
			p.width = msg.width - 2
			p.height = msg.height

			return p, nil
		}
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			p.OnSelect(GameStateGlobal.ShipManager.Ships[p.currentShipID])
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.sortedShips)-1 {
				p.cursor++
			}
		case "esc":
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *ShipManagementPane) View() string {

	var rows []string

	idx := 0
	for _, ship := range p.sortedShips {
		var row string
		if idx == p.cursor {
			p.currentShipID = ship.GetID()

			row = fmt.Sprintf("%s - %s", ship.GetName(), ship.GetPosition())
			row = Style.Border(lipgloss.NormalBorder()).Padding(0, 1).Render(row)
			row = Theme.focusedStyle.Render(row)
		} else {
			row = fmt.Sprintf("%s - %s", ship.GetName(), ship.GetPosition())
			row = Style.Border(lipgloss.NormalBorder()).Padding(0, 1).Render(row)
			row = Theme.blurredStyle.Render(row)
		}

		rows = append(rows, row)
		idx++
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	title := "Ships"
	title = Style.Width(p.width).AlignHorizontal(lipgloss.Center).Render(title)

	content = lipgloss.JoinVertical(lipgloss.Left, title, content)

	return content

}
