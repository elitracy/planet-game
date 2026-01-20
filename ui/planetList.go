package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
)

type PlanetListPane struct {
	*Pane

	planets []*models.Planet
	cursor  int
	theme   UITheme
}

func NewPlanetListPane(planets []*models.Planet, title string) *PlanetListPane {
	return &PlanetListPane{
		Pane: &Pane{
			title: title,
		},
		planets: planets,
	}
}

func (p *PlanetListPane) Init() tea.Cmd {
	pane := &PlanetInfoPane{
		Pane: &Pane{
			title: p.planets[p.cursor].Name,
		},
		planet: p.planets[p.cursor],
	}
	paneID := PaneManager.AddPane(pane)
	return pushDetailStackCmd(paneID)
}

func (p *PlanetListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		p.width = msg.width
		p.height = msg.height
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			} else {
				p.cursor = len(p.planets) - 1
			}
			pane := &PlanetInfoPane{
				Pane: &Pane{
					title: p.planets[p.cursor].Name,
				},
				planet: p.planets[p.cursor],
			}
			paneID := PaneManager.AddPane(pane)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case "down", "j":
			if p.cursor < len(p.planets)-1 {
				p.cursor++
			} else {
				p.cursor = 0
			}

			pane := &PlanetInfoPane{
				Pane: &Pane{
					title: p.planets[p.cursor].Name,
				},
				planet: p.planets[p.cursor],
			}
			paneID := PaneManager.AddPane(pane)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case "esc":
			return p, tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
		case "enter":
			pane := &PlanetInfoPane{
				Pane: &Pane{
					title: p.planets[p.cursor].Name,
				},
				planet: p.planets[p.cursor],
			}
			paneID := PaneManager.AddPane(pane)
			return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *PlanetListPane) View() string {
	p.theme = GetPaneTheme(p)

	s := "Available Planets:\n"

	for i, choice := range p.planets {
		cursor := " "
		if p.cursor == i {
			cursor = ">"
			s += p.theme.FocusedStyle.Render(fmt.Sprintf("%s %s", cursor, choice.Name))
		} else {
			s += fmt.Sprintf("%s %s", cursor, choice.Name)
		}

		if choice.ColonyName != "" {
			colony := fmt.Sprintf(" (%s)", choice.ColonyName)
			colony = Style.Foreground(lipgloss.Color("240")).Render(colony)
			s += colony
		}

		s += "\n"
	}
	return s
}
