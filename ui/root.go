package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	. "github.com/elitracy/planets/core/interfaces"
)

type RootPane struct {
	id     int
	title  string
	width  int
	height int

	tabCursor int
	focusTabs bool
	tabs   []int // paneIDs -> TODO: make PaneID type
	activePane Pane
}

func NewRootPane(title string, tabs []int) *RootPane {
	pane := &RootPane{
		title: title,
		tabs:  tabs,
		focusTabs: true,
	}

	id := tabs[pane.tabCursor]
	pane.activePane = PaneManager.Panes[id].(Pane)

	return pane
}

func (p RootPane) GetId() int       { return p.id }
func (p *RootPane) SetId(id int)    { p.id = id }
func (p RootPane) GetTitle() string { return p.title }
func (p RootPane) GetWidth() int    { return p.width }
func (p RootPane) GetHeight() int   { return p.height }
func (p *RootPane) SetWidth(w int)  { p.width = w }
func (p *RootPane) SetHeight(h int) { p.height = h }

func (p *RootPane) Init() tea.Cmd { return nil }

func (p *RootPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
			if p.focusTabs && p.tabCursor > 0 {
				p.tabCursor--
			}
		case "down", "j":
			if p.focusTabs && p.tabCursor < len(p.tabs)-1 {
				p.tabCursor++
			}else{
			}
		case "enter":
			id :=p.tabs[p.tabCursor]
			pane:= PaneManager.Panes[id]

			p.activePane = pane.(Pane)

			p.focusTabs = false
		case "esc":
			p.focusTabs = true 
			return p, nil
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

	content = lipgloss.JoinHorizontal(lipgloss.Left, content, p.activePane.(tea.Model).View())

	content = consts.Style.Render(content)
	return content
}
