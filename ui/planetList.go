package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
)

type PlanetList struct {
	choices  []*models.Planet
	cursor   int
	selected int
	id       int
	title    string
}

func NewPlanetList(planets []*models.Planet, title string) *PlanetList {
	return &PlanetList{
		choices:  planets,
		selected: -1,
		title:    title,
	}
}

func (p PlanetList) GetId() int       { return p.id }
func (p *PlanetList) SetId(id int)    { p.id = id }
func (p PlanetList) GetTitle() string { return p.title }

func (p *PlanetList) Init() tea.Cmd {
	return nil
}

func (p *PlanetList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var childPaneID int

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.choices)-1 {
				p.cursor++
			}
		case " ":
			p.selected = p.cursor

			pane := &PlanetInfoPane{
				id:     0,
				title:  "Planet Info",
				planet: p.choices[p.cursor],
			}
			childPaneID := PaneManager.AddPane(pane)

			return p, pushFocusCmd(childPaneID)

		case "esc":
			PaneManager.RemovePane(childPaneID)
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *PlanetList) View() string {
	s := "Available Planets:\n"

	for i, choice := range p.choices {
		cursor := " "
		if p.cursor == i {
			cursor = ">"
			s += Theme.focusedStyle.Render(fmt.Sprintf("%s %s", cursor, choice.Name))
		} else {
			s += fmt.Sprintf("%s %s", cursor, choice.Name)
		}

		if choice.ColonyName != "" {
			colony := fmt.Sprintf(" (%s)", choice.ColonyName)
			colony = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(colony)
			s += colony
		}

		s += "\n"
	}
	return s
}

// planet dashboard

// display telemetry

// manage contructions
// build more of each

// upgrade each
