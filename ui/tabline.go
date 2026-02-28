package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
)

type TabLinePane struct {
	*engine.Pane

	cursor    int
	tabs      []*engine.LayoutNode
	tabTitles []string
	theme     UITheme
}

func NewTablinePane() *TabLinePane {
	pane := &TabLinePane{
		Pane: engine.NewPane("Tabs", nil),
		tabTitles: []string{
			"Systems",
			"Orders",
			"Messages",
		},
	}

	return pane
}

func (p *TabLinePane) Init() tea.Cmd { return nil }

func (p *TabLinePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return p, tea.Quit
		case "shift+tab":
			if p.cursor > 0 {
				p.cursor--
			} else {
				p.cursor = len(p.tabs) - 1
			}
			layout := p.tabs[p.cursor]

			return p, tea.Sequence(
				setLayoutCmd(layout),
				flushFocusStackCmd(),
				pushFocusStackCmd(layout.Pane.ID()),
			)
		case "tab":
			if p.cursor < len(p.tabs)-1 {
				p.cursor++
			} else {
				p.cursor = 0
			}
			layout := p.tabs[p.cursor]

			return p, tea.Sequence(
				setLayoutCmd(layout),
				flushFocusStackCmd(),
				pushFocusStackCmd(layout.Pane.ID()),
			)
		}

	}
	return p, nil
}

func (p *TabLinePane) View() string {
	p.theme = Theme

	title := "Tabs: "
	var tabs []string
	for i := range p.tabs {
		tabTitle := fmt.Sprintf("[%v] ", p.tabTitles[i])
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
