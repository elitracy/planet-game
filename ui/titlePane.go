package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/logging"
)

type TitlePane struct {
	BasePane
}

func NewTitlePane(text string, id int) *TitlePane {
	return &TitlePane{BasePane{title: text, id: id}}
}

func (p TitlePane) Init() tea.Cmd { return tick(p.id) }
func (p TitlePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		if msg.id == p.id {
			logging.Log(fmt.Sprintf("Pane: %d @ %d", p.id, msg.id), "UI")
			return p, tick(p.id)
		}
	case tea.KeyMsg:
		logging.Log("Key Pressed: "+msg.String(), "UI")
	}
	return p, nil
}

type tickMsg struct{ id int }

func tick(id int) tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg { return tickMsg{id} })
}

func (p TitlePane) View() string { return p.title }

func (p TitlePane) GetId() int { return p.id }
