package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	. "github.com/elitracy/planets/core/interfaces"
)

type TabSelectPane struct {
	id     int
	title  string
	width  int
	height int

	cursor int
	tabs   []int // paneIDs -> TODO: make PaneID type
}

func NewTabSelectPane(title string, tabs []int) *TabSelectPane {
	pane := &TabSelectPane{
		title: title,
		tabs:  tabs,
	}

	return pane
}

func (p TabSelectPane) GetId() int       { return p.id }
func (p *TabSelectPane) SetId(id int)    { p.id = id }
func (p TabSelectPane) GetTitle() string { return p.title }
func (p TabSelectPane) GetWidth() int    { return p.width }
func (p TabSelectPane) GetHeight() int   { return p.height }
func (p *TabSelectPane) SetWidth(w int)  { p.width = w }
func (p *TabSelectPane) SetHeight(h int) { p.height = h }

func (p *TabSelectPane) Init() tea.Cmd { return nil }

func (p *TabSelectPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.GetId() {
			p.width = int(float32(msg.width) * 0.1)
			p.height = msg.height

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
			}
		case "esc":
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *TabSelectPane) View() string {

	title := consts.Style.Width(p.width).AlignHorizontal(lipgloss.Center).Render(p.title)

	rows := []string{title}
	for i, id := range p.tabs {
		if pane, ok := PaneManager.Panes[id].(Pane); ok {

			var row string
			if i == p.cursor {
				row = consts.Theme.FocusedStyle.Padding(0, 1).Render(pane.GetTitle())
			} else {
				row = consts.Theme.BlurredStyle.Padding(0, 1).Render(pane.GetTitle())
			}

			rows = append(rows, row)
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return content
}
