package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
)

type RootPane struct {
	*Pane

	focusTabs      bool
	cursor         int
	tabs           []core.PaneID
	activePane     ManagedPane
	lastActivePane ManagedPane
}

func NewRootPane(title string, tabs []core.PaneID) *RootPane {
	pane := &RootPane{
		Pane: &Pane{
			title: title,
		},
		tabs:      tabs,
		focusTabs: true,
	}

	id := tabs[pane.cursor]
	pane.activePane = PaneManager.Panes[id]

	return pane
}

func (p *RootPane) Init() tea.Cmd {
	p.lastActivePane = PaneManager.Panes[p.tabs[0]]
	return nil
}

func (p *RootPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case focusTabsMsg:
		p.focusTabs = true
		if msg.lastActiveID == -1 {
			p.lastActivePane = PaneManager.Panes[p.tabs[0]]
		} else {
			p.lastActivePane = PaneManager.Panes[msg.lastActiveID]
			logging.Info("Last Active: %v", p.lastActivePane.Title())
		}
	case paneResizeMsg:
		if msg.paneID == p.Pane.id {
			p.Pane.width = int(float32(msg.width) * 0.15)
			p.Pane.height = msg.height

			return p, nil
		}
	case core.TickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.tabs)-1 {
				p.cursor++
			} else {
			}
		case "enter":
			id := p.tabs[p.cursor]
			pane := PaneManager.Panes[id]

			p.activePane = pane

			p.focusTabs = false
			return p, pushFocusCmd(id)
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *RootPane) View() string {

	title := consts.Style.Width(p.Pane.width).AlignHorizontal(lipgloss.Center).Render(p.Pane.Title())

	rows := []string{title}
	for i, paneID := range p.tabs {
		pane := PaneManager.Panes[paneID]

		var row string
		if i == p.cursor {
			row = consts.Theme.FocusedStyle.Padding(0, 1).Render(pane.Title())
			if !p.focusTabs {
				row = consts.Theme.DimmedStyle.Padding(0, 1).Render(pane.Title())
			}
		} else {
			row = consts.Theme.BlurredStyle.Padding(0, 1).Render(pane.Title())
		}

		rows = append(rows, row)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	content = consts.Style.
		Height(p.Pane.height).
		Padding(1).
		Border(lipgloss.ThickBorder(), false, true, false, false).
		MarginRight(3).
		Render(content)

	if p.focusTabs {
		content = lipgloss.JoinHorizontal(lipgloss.Left, content, p.lastActivePane.View())
	} else {
		content = lipgloss.JoinHorizontal(lipgloss.Left, content, PaneManager.ActivePane().View())
	}

	content = consts.Style.Render(content)

	return content
}
