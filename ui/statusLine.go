package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/config"
)

type StatusLinePane struct {
	*engine.Pane
}

func NewStatusLinePane(tick engine.Tick) *StatusLinePane {
	return &StatusLinePane{Pane: engine.NewPane("Status Line", nil)}
}

func (p *StatusLinePane) Init() tea.Cmd { return nil }

func (p *StatusLinePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

func (p *StatusLinePane) View() string {
	focusedPane := PaneManager.Panes[PaneManager.PeekFocusStack()]

	content := ""

	components := strings.Split(config.FormatGameTime(game.State.CurrentTick), ".")
	components[2] = Theme.DimmedStyle.Render(components[2])
	componentsStyled := strings.Join(components, ".")

	content += fmt.Sprintf("Time : %v", componentsStyled)

	var keys *engine.KeyBindings
	if focusedPane != nil {
		keys = focusedPane.GetKeys()
	}

	keysStyled := ""
	if keys != nil {
		keysStyled = Theme.DimmedStyle.Render(keys.String())
	}

	content = lipgloss.JoinVertical(lipgloss.Left, content, keysStyled)

	return content
}

func (p *StatusLinePane) SetKeys(keys *engine.KeyBindings) { p.Pane.SetKeys(keys) }
