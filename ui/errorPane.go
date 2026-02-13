package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/engine"
)

type ErrorPane struct {
	*engine.Pane

	errorMsg string
}

func NewErrorPane(errorMsg string) *ErrorPane {
	return &ErrorPane{
		Pane:     &engine.Pane{},
		errorMsg: errorMsg,
	}
}

func (p *ErrorPane) Init() tea.Cmd                       { return nil }
func (p *ErrorPane) Update(tea.Msg) (tea.Model, tea.Cmd) { return p, nil }
func (p *ErrorPane) View() string                        { return p.errorMsg }
