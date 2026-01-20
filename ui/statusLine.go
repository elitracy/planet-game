package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
)

type StatusLinePane struct {
	*Pane
	currentCycle core.Cycle
}

func NewStatusLinePane(tick core.Tick) *StatusLinePane {
	return &StatusLinePane{Pane: &Pane{}, currentCycle: core.TickToCycle(tick)}
}

func (p *StatusLinePane) Init() tea.Cmd {
	return nil
}

func (p *StatusLinePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case core.TickMsg:
		logging.Info("tick msg: %v", msg)
		p.currentCycle = core.TickToCycle(msg.Tick)
	}

	return p, nil
}

func (p *StatusLinePane) View() string {
	content := ""

	content += fmt.Sprintf("Current Cycle: %3.4f", p.currentCycle)

	return content
}
