package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/models"
)

type planetListModel struct {
	choices  []*models.Planet
	cursor   int
	selected map[int]struct{}
}

func CreatePlanetListInitialModel(planets []*models.Planet) planetListModel {
	return planetListModel{
		choices:  planets,
		selected: make(map[int]struct{}),
	}
}

func (m planetListModel) Init() tea.Cmd {
	return nil
}

func (m planetListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m planetListModel) View() string {
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

	s += "\nPress q or C-c to quit.\n"
	return s
}

// planet dashboard

// display telemetry

// manage contructions
// build more of each

// upgrade each
