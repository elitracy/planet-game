package ui

import tea "github.com/charmbracelet/bubbletea"

type ErrorPane struct {
	*Pane

	errorMsg string
}

func NewErrorPane(errorMsg string) *ErrorPane {
	return &ErrorPane{
		Pane: &Pane{},
		errorMsg: errorMsg,
	}
}

func (p *ErrorPane) Init() tea.Cmd                       { return nil }
func (p *ErrorPane) Update(tea.Msg) (tea.Model, tea.Cmd) { return p, nil }
func (p *ErrorPane) View() string                        { return p.errorMsg }
