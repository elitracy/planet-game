package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/interfaces"
	. "github.com/elitracy/planets/core/interfaces"
	"github.com/elitracy/planets/core/logging"
)

type RootPane struct {
	id     core.PaneID
	title  string
	width  int
	height int

	focusTabs      bool
	tabCursor      int
	tabs           []core.PaneID
	activePane     tea.Model
	lastActivePane tea.Model
}

func NewRootPane(title string, tabs []core.PaneID) *RootPane {
	pane := &RootPane{
		title:     title,
		tabs:      tabs,
		focusTabs: true,
	}

	id := tabs[pane.tabCursor]
	pane.activePane = PaneManager.Panes[id]

	return pane
}

func (p RootPane) GetId() core.PaneID    { return p.id }
func (p *RootPane) SetId(id core.PaneID) { p.id = id }
func (p RootPane) GetTitle() string      { return p.title }
func (p RootPane) GetWidth() int         { return p.width }
func (p RootPane) GetHeight() int        { return p.height }
func (p *RootPane) SetWidth(w int)       { p.width = w }
func (p *RootPane) SetHeight(h int)      { p.height = h }

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
			logging.Info("Last Active: %v", p.lastActivePane.(interfaces.Pane).GetTitle())
		}
	case paneResizeMsg:
		if msg.paneID == p.GetId() {
			p.width = int(float32(msg.width) * 0.15)
			p.height = msg.height

			return p, nil
		}
	case core.TickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.tabCursor > 0 {
				p.tabCursor--
			}
		case "down", "j":
			if p.tabCursor < len(p.tabs)-1 {
				p.tabCursor++
			} else {
			}
		case "enter":
			id := p.tabs[p.tabCursor]
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

	title := consts.Style.Width(p.width).AlignHorizontal(lipgloss.Center).Render(p.title)

	rows := []string{title}
	for i, id := range p.tabs {

		if pane, ok := PaneManager.Panes[id].(Pane); ok {

			var row string
			if i == p.tabCursor {
				row = consts.Theme.FocusedStyle.Padding(0, 1).Render(pane.GetTitle())
			} else {
				row = consts.Theme.BlurredStyle.Padding(0, 1).Render(pane.GetTitle())
			}

			rows = append(rows, row)
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	content = consts.Style.
		Height(p.height).
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
