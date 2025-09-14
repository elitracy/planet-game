package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TitlePane struct {
	BasePane
	id    int
	title string
}

func NewTitlePane(text string, id int) *TitlePane {
	return &TitlePane{title: text, id: id}
}

func (p *TitlePane) Init() tea.Cmd { return tick(p.id) }
func (p *TitlePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		if msg.id == p.id {
			return p, tick(p.id)
		}
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

type tickMsg struct{ id int }

func tick(id int) tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg { return tickMsg{id} })
}

func (p *TitlePane) View() string { return p.title }

func (p TitlePane) GetId() int       { return p.id }
func (p TitlePane) GetTitle() string { return p.title }
