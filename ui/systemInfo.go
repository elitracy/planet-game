package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
)

type StarSystemInfoPane struct {
	*Pane

	cursor int
	system *models.StarSystem
	theme  UITheme
}

func (p *StarSystemInfoPane) Init() tea.Cmd { return nil }
func (p *StarSystemInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case "enter":
			pane := &PlanetInfoPane{
				Pane: &Pane{
					title: "Planet Info",
				},
				planet: p.system.Planets[p.cursor],
			}
			paneID := PaneManager.AddPane(pane)
			return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
		case "ctrl+c", "q":
			return p, tea.Quit
		case "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "j":
			if p.cursor < len(p.system.Planets)-1 {
				p.cursor++
			}
		}

	default:
	}
	return p, nil
}

func (p *StarSystemInfoPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.system.Name

	var rows []string
	for i, planet := range p.system.Planets {
		row := planet.Name

		if planet.ColonyName != "" {
			row += fmt.Sprintf(" [%v]", planet.ColonyName)
		}

		row += fmt.Sprintf(" %v - population %v", planet.Position, planet.Population)

		if i == p.cursor {
			row = p.theme.FocusedStyle.Render(row)
		}

		rows = append(rows, row)
	}

	planetList := lipgloss.JoinVertical(lipgloss.Left, rows...)
	planetList = Style.Padding(0, 1).Border(lipgloss.RoundedBorder(), true, false, false, false).Render(planetList)

	content := lipgloss.JoinVertical(lipgloss.Left, title, planetList)
	return content

}

func NewSystemInfoPane(title string, system *models.StarSystem) *StarSystemInfoPane {
	return &StarSystemInfoPane{
		Pane: &Pane{
			title: title,
		},
		system: system,
	}
}
