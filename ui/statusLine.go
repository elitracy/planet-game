package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
)

type StatusLinePane struct {
	*Pane
}

func NewStatusLinePane(tick core.Tick) *StatusLinePane {
	return &StatusLinePane{Pane: &Pane{}}
}

func (p *StatusLinePane) Init() tea.Cmd {
	return nil
}

func (p *StatusLinePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

func (p *StatusLinePane) View() string {
	content := ""

	components := strings.Split(state.State.Tick.String(), ".")
	components[2] = Theme.DimmedStyle.Render(components[2])
	componentsStyled := strings.Join(components, ".")

	content += fmt.Sprintf("Time : %v", componentsStyled)

	return content
}
