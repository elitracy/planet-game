package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/models"
)

type PlanetList struct {
	choices  []*models.Planet
	cursor   int
	selected map[int]struct{}
	id       int
	title    string
}

func NewPlanetList(planets []*models.Planet, id int, title string) PlanetList {
	return PlanetList{
		choices:  planets,
		selected: make(map[int]struct{}),
		id:       id,
		title:    title,
	}
}

func (p PlanetList) GetId() int {
	return p.id
}

func (p PlanetList) GetTitle() string {
	return p.title
}

func (m PlanetList) Init() tea.Cmd {
	return nil
}

func (m PlanetList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "esc":
			return PopFocus(), nil
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m PlanetList) View() string {
	s := "Available Planets:\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}
	return s
}

// planet dashboard

// display telemetry

// manage contructions
// build more of each

// upgrade each
