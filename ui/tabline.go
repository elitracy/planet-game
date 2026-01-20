package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core/consts"
)

type TablinePane struct {
	*Pane

	cursor int
	tabs   []ManagedPane
}

func NewTablinePane(title string, tabs []ManagedPane) *TablinePane {
	pane := &TablinePane{
		Pane: &Pane{
			title: title,
		},
		tabs: tabs,
	}

	return pane
}

func (p *TablinePane) Init() tea.Cmd {
	return nil
}

func (p *TablinePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.Pane.id {
			p.Pane.width = msg.width - 2
			p.Pane.height = msg.height

			return p, nil
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popMainFocusCmd(p.Pane.id)
		case "ctrl+c", "q":
			return p, tea.Quit
		case "shift+tab":
			if p.cursor > 0 {
				p.cursor--
				PaneManager.SetMainPane(p.tabs[p.cursor])
			}
			flushDetailStackCmd()
		case "tab":
			if p.cursor < len(p.tabs)-1 {
				p.cursor++
				PaneManager.SetMainPane(p.tabs[p.cursor])
			}
			flushDetailStackCmd()
		}

	}
	return p, tea.Batch(cmds...)
}

func (p *TablinePane) View() string {

	title := "Tabs: "
	var tabs []string
	for i, tab := range p.tabs {
		tabTitle := fmt.Sprintf("[%v] ", tab.Title())
		if p.cursor == i {
			tabs = append(tabs, consts.Theme.FocusedStyle.Render(tabTitle))
		} else {
			tabs = append(tabs, consts.Theme.BlurredStyle.Render(tabTitle))
		}
	}
	tabContent := lipgloss.JoinHorizontal(lipgloss.Left, tabs...)
	content := lipgloss.JoinHorizontal(lipgloss.Left, title, tabContent)
	return content
}
