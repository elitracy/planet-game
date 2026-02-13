package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
)

type InfoTablePane struct {
	*engine.Pane
	table   table.Model
	keymaps map[string]func() tea.Cmd
	theme   UITheme
}

func NewInfoTablePane(table table.Model, keymaps map[string]func() tea.Cmd) *InfoTablePane {
	pane := &InfoTablePane{
		Pane:    engine.NewPane("Info Table", nil),
		table:   table,
		keymaps: keymaps,
	}
	return pane
}
func (p *InfoTablePane) Init() tea.Cmd {
	return nil
}

func (p *InfoTablePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if handler, ok := p.keymaps[key]; ok {
			return p, handler()
		}

		switch msg.String() {
		case "k", "up":
			p.table.MoveUp(1)
		case "j", "down":
			p.table.MoveDown(1)
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	default:
		model, cmd := p.table.Update(msg)
		p.table = model
		return p, cmd
	}
	return p, nil

}

func (p *InfoTablePane) View() string {
	tableStyles := table.DefaultStyles()

	tableStyles.Header = tableStyles.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(p.theme.DimmedStyle.GetForeground()).
		BorderBottom(true).
		Bold(false)

	tableStyles.Selected = tableStyles.Selected.
		Foreground(p.theme.FocusedStyle.GetBackground()).
		Background(p.theme.FocusedStyle.GetForeground()).
		Bold(false)

	p.table.SetStyles(tableStyles)

	return p.table.View()
}

func (p *InfoTablePane) SetTheme(theme UITheme) {
	p.theme = theme
}
