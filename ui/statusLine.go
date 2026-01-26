package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/state"
)

type StatusLinePane struct {
	*Pane
	keys *string
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

	if len(*p.keys) == 0 {
		return content
	}

	keysStyled := Theme.DimmedStyle.Render(*p.keys)
	content = lipgloss.JoinVertical(lipgloss.Left, content, keysStyled)

	return content
}

func (p *StatusLinePane) SetKeys(keys *string) { p.keys = keys }
