package ui

import tea "github.com/charmbracelet/bubbletea"

type pane struct {
	id    int
	label string
}

func (p pane) Init() tea.Cmd                           { return nil }
func (p pane) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }
func (p pane) View() string                            { return p.label }

type dashboard struct {
	grid[][]
} sC
