package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	. "github.com/elitracy/planets/models"
)

type TitlePane struct {
	Pane
	id    int
	title string
}

func NewTitlePane(text string) *TitlePane {
	return &TitlePane{title: text}
}

func (p *TitlePane) Init() tea.Cmd { return nil }
func (p *TitlePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p *TitlePane) View() string { return p.title }

func (p TitlePane) GetId() int       { return p.id }
func (p *TitlePane) SetId(id int)    { p.id = id }
func (p TitlePane) GetTitle() string { return p.title }
