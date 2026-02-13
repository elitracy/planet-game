package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
)

type TabLinePane struct {
	*engine.Pane

	cursor int
	tabs   []engine.ManagedPane
	theme  UITheme
}

func NewTablinePane(tabs []engine.ManagedPane) *TabLinePane {
	pane := &TabLinePane{
		Pane: engine.NewPane("Tabs", nil),
		tabs: tabs,
	}

	return pane
}

func (p *TabLinePane) Init() tea.Cmd { return nil }

func (p *TabLinePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.Pane.ID() {
			p.SetSize(msg.width-2, msg.height)

			return p, nil
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popMainFocusCmd(p.Pane.ID())
		case "ctrl+c", "q":
			return p, tea.Quit
		case "shift+tab":
			if p.cursor > 0 {
				p.cursor--
			} else {
				p.cursor = len(p.tabs) - 1
			}
			pane := p.tabs[p.cursor]

			return p, tea.Sequence(
				flushDetailStackCmd(),
				flushFocusStackCmd(),
				setMainFocusCmd(pane.ID()),
				pushFocusStackCmd(pane.ID()),
			)
		case "tab":
			if p.cursor < len(p.tabs)-1 {
				p.cursor++
			} else {
				p.cursor = 0
			}
			pane := p.tabs[p.cursor]

			return p, tea.Sequence(
				flushDetailStackCmd(),
				flushFocusStackCmd(),
				setMainFocusCmd(pane.ID()),
				pushFocusStackCmd(pane.ID()),
			)
		}

	}
	return p, nil
}

func (p *TabLinePane) View() string {
	p.theme = Theme

	title := "Tabs: "
	var tabs []string
	for i, tab := range p.tabs {
		tabTitle := fmt.Sprintf("[%v] ", tab.Title())
		if p.cursor == i {
			tabs = append(tabs, p.theme.FocusedStyle.Render(tabTitle))
		} else {
			tabs = append(tabs, p.theme.BlurredStyle.Render(tabTitle))
		}
	}
	tabContent := lipgloss.JoinHorizontal(lipgloss.Left, tabs...)
	content := lipgloss.JoinHorizontal(lipgloss.Left, title, tabContent)
	return content
}
