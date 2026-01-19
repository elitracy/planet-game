package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
)

type PlanetListPane struct {
	*Pane

	planets []*models.Planet
	cursor  int
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
	return nil
}

func (p *PlanetListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var childPaneID core.PaneID

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.planets)-1 {
				p.cursor++
			}
		case "enter":
			pane := &PlanetInfoPane{
				Pane: &Pane{
					title: "Planet Info",
				},
				planet: p.planets[p.cursor],
			}
			childPaneID := PaneManager.AddPane(pane)
			return p, pushFocusCmd(childPaneID)

		case "esc":
			PaneManager.RemovePane(childPaneID)
			return p, popFocusCmd(p.Pane.id)
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *PlanetListPane) View() string {
	s := "Available Planets:\n"

	for i, choice := range p.planets {
		cursor := " "
		if p.cursor == i && PaneManager.ActivePane().ID() == p.Pane.ID() {
			cursor = ">"
			s += consts.Theme.FocusedStyle.Render(fmt.Sprintf("%s %s", cursor, choice.Name))
		} else {
			s += fmt.Sprintf("%s %s", cursor, choice.Name)
		}

		if choice.ColonyName != "" {
			colony := fmt.Sprintf(" (%s)", choice.ColonyName)
			colony = consts.Style.Foreground(lipgloss.Color("240")).Render(colony)
			s += colony
		}

		s += "\n"
	}
	return s
}
