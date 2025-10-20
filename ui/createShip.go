package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
)

var ships = []string{
	"Scout - recon & intelligence",
	"Fighter - attack & defense",
	"Cargo - transport resources",
}

type CreateShipPane struct {
	id        int
	title     string
	width     int
	cursor    int
	ShipTypes models.ShipType
}

func (p CreateShipPane) GetId() int       { return p.id }
func (p *CreateShipPane) SetId(id int)    { p.id = id }
func (p CreateShipPane) GetTitle() string { return p.title }

func (p *CreateShipPane) Init() tea.Cmd { return nil }
func (p *CreateShipPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.GetId() {
			p.width = msg.width - 2
			return p, nil
		}
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			order := NewBuildShipOrder()
		case "j":
			if p.cursor < len(ships)-1 {
				p.cursor++
			}
		case "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "esc":
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	default:
	}
	return p, nil
}

func (p *CreateShipPane) View() string {
	title := p.title
	title = Style.Width(p.width).AlignHorizontal(lipgloss.Center).Render(title)

	rows := []string{title}

	for i, ship := range ships {

		row := Style.Border(lipgloss.NormalBorder()).Padding(0, 1).Render(ship)
		if p.cursor == i {
			row = Theme.focusedStyle.Render(ship)
		} else {
			row = Theme.blurredStyle.Render(ship)
		}

		rows = append(rows, row)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return content

}

func NewCreateShipPane(title string) *CreateShipPane {
	pane := &CreateShipPane{
		title: title,
	}

	return pane
}
