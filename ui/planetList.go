package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/interfaces"
	"github.com/elitracy/planets/models"
)

type PlanetListPane struct {
	id     core.PaneID
	title  string
	width  int
	height int

	choices []*models.Planet
	cursor  int
}

func (p PlanetListPane) GetId() core.PaneID    { return p.id }
func (p *PlanetListPane) SetId(id core.PaneID) { p.id = id }
func (p PlanetListPane) GetTitle() string      { return p.title }
func (p PlanetListPane) GetWidth() int         { return p.width }
func (p PlanetListPane) GetHeight() int        { return p.height }
func (p *PlanetListPane) SetWidth(w int)       { p.width = w }
func (p *PlanetListPane) SetHeight(h int)      { p.height = h }

func NewPlanetListPane(planets []*models.Planet, title string) *PlanetListPane {
	return &PlanetListPane{
		choices: planets,
		title:   title,
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
			if p.cursor < len(p.choices)-1 {
				p.cursor++
			}
		case "enter":
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

func (p *PlanetListPane) View() string {
	s := "Available Planets:\n"

	for i, choice := range p.choices {
		cursor := " "
		if p.cursor == i && PaneManager.ActivePane().(interfaces.Pane).GetId() == p.GetId() {
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

// planet dashboard

// display telemetry

// manage contructions
// build more of each

// upgrade each
