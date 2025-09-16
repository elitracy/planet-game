package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TitlePane struct {
	Pane
	id    int
	title string
}

func NewTitlePane(text string, id int) *TitlePane {
	return &TitlePane{title: text, id: id}
}

func (p *TitlePane) Init() tea.Cmd { return tick() }
func (p *TitlePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		return p, tick()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return PopFocus(), nil
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *TitlePane) View() string { return p.title }

func (p TitlePane) GetId() int       { return p.id }
func (p TitlePane) GetTitle() string { return p.title }
