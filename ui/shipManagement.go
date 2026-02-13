package ui

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/models"
)

type ShipManagementPane struct {
	*engine.Pane
	cursor        int
	currentShipID int
	sortedShips   []*models.Ship
	manager       *models.ShipManager
	OnSelect      func(ship *models.Ship)
	theme         UITheme
}

func CreateNewShipManagementPane(title string, shipManager *models.ShipManager, callback func(ship *models.Ship)) *ShipManagementPane {
	pane := &ShipManagementPane{
		Pane:     engine.NewPane(title, engine.NewKeyBindings()),
		manager:  shipManager,
		OnSelect: callback,
	}

	return pane
}

func (p *ShipManagementPane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Select, "enter").
		Set(engine.Back, "esc").
		Set(engine.Up, "k").
		Set(engine.Down, "j").
		Set(engine.Quit, "q")

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
		if msg.paneID == p.Pane.ID() {
			p.SetSize(msg.width-2, msg.height)

			return p, nil
		}
	case tea.KeyMsg:
		switch msg.String() {
		case p.GetKeys().Get(engine.Select):
			p.OnSelect(game.State.ShipManager.Ships[p.currentShipID])
		case p.GetKeys().Get(engine.Up):
			if p.cursor > 0 {
				p.cursor--
			}
		case p.GetKeys().Get(engine.Down):
			if p.cursor < len(p.sortedShips)-1 {
				p.cursor++
			}
		case p.GetKeys().Get(engine.Back):
			return p, tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
		case p.GetKeys().Get(engine.Quit):
			return p, tea.Quit
		}
	}

	p.currentShipID = p.sortedShips[p.cursor].GetID()

	return p, nil
}

func (p *ShipManagementPane) View() string {
	p.theme = GetPaneTheme(p)

	var rows []string

	idx := 0
	for _, ship := range p.sortedShips {
		var row string
		if idx == p.cursor {

			row = fmt.Sprintf("%s - %s", ship.GetName(), ship.GetLocation())
			row = Style.Border(lipgloss.NormalBorder()).Padding(0, 1).Render(row)
			row = p.theme.FocusedStyle.Render(row)
		} else {
			row = fmt.Sprintf("%s - %s", ship.GetName(), ship.GetLocation())
			row = Style.Border(lipgloss.NormalBorder()).Padding(0, 1).Render(row)
			row = p.theme.BlurredStyle.Render(row)
		}

		rows = append(rows, row)
		idx++
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	title := "Ships"
	title = Style.Width(p.Pane.Width()).AlignHorizontal(lipgloss.Center).Render(title)

	content = lipgloss.JoinVertical(lipgloss.Left, title, content)

	return content

}
