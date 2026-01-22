package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TitlePane struct {
	*Pane
}

func NewTitlePane(text string) *TitlePane {
	return &TitlePane{
		Pane: &Pane{
			title: text,
		},
	}
}

func (p *TitlePane) Init() tea.Cmd { return nil }
func (p *TitlePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popMainFocusCmd(p.ID())
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *TitlePane) View() string { return p.title }
