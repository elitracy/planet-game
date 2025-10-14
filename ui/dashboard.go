package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dashboard struct {
	Grid      [][]int
	ActiveRow int
	ActiveCol int

	id    int
	title string
}

const (
	TERM_PADDING = 2
)

func (p Dashboard) GetId() int       { return p.id }
func (p *Dashboard) SetId(id int)    { p.id = id }
func (p Dashboard) GetTitle() string { return p.title }

func (p *Dashboard) Init() tea.Cmd {
	cmds := p.getPaneResizeCommands()
	return tea.Batch(cmds...)
}

func (p *Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmds := p.getPaneResizeCommands()
		return p, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmds := p.getPaneResizeCommands()
			cmds = append(cmds, pushFocusCmd(p.Grid[p.ActiveRow][p.ActiveCol]))

			return p, tea.Batch(cmds...)
		case "esc":
			cmds := p.getPaneResizeCommands()
			cmds = append(cmds, popFocusCmd())
			return p, tea.Batch(cmds...)
		case "h":
			if p.ActiveCol > 0 {
				p.ActiveCol--
			}
		case "l":
			if p.ActiveCol < len(p.Grid[p.ActiveRow])-1 {
				p.ActiveCol++
			}
		case "k":
			if p.ActiveRow > 0 {
				if p.ActiveCol >= len(p.Grid[p.ActiveRow-1]) {
					p.ActiveCol = len(p.Grid[p.ActiveRow-1]) - 1
				}

				p.ActiveRow--
			}
		case "j":
			if p.ActiveRow < len(p.Grid)-1 {
				if p.ActiveCol >= len(p.Grid[p.ActiveRow+1]) {
					p.ActiveCol = len(p.Grid[p.ActiveRow+1]) - 1
				}
				p.ActiveRow++
			}
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	}

	return p, nil
}

func (p *Dashboard) View() string {

	render := make([][]string, len(p.Grid))
	for r := range p.Grid {
		render[r] = make([]string, len(p.Grid[r]))
		cols := len(p.Grid[r])
		paneWidth := PaneManager.Width / cols
		paneHeight := PaneManager.Height / len(p.Grid)
		activeStyle := Style.
			Width(paneWidth).
			Height(paneHeight).
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("212"))

		inactiveStyle := Style.
			Width(paneWidth).
			Height(paneHeight).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

		for c := range p.Grid[r] {
			paneID := p.Grid[r][c]
			if r == p.ActiveRow && c == p.ActiveCol {
				render[r][c] = activeStyle.Render(PaneManager.Panes[paneID].View())
			} else {
				render[r][c] = inactiveStyle.Render(PaneManager.Panes[paneID].View())
			}

		}
	}

	var rows []string
	for r := range render {
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, render[r]...))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	content = Style.Render(content)
	return content
}

func NewDashboard(grid [][]int, activeRow, activeCol int, title string) *Dashboard {
	return &Dashboard{
		Grid:      grid,
		ActiveRow: activeRow,
		ActiveCol: activeCol,
		title:     title,
	}
}

func (p Dashboard) getPaneResizeCommands() []tea.Cmd {
	var cmds []tea.Cmd
	for r := range p.Grid {
		cols := len(p.Grid[r])
		for c := range p.Grid[r] {
			paneWidth := PaneManager.Width / cols
			paneHeight := PaneManager.Height / len(p.Grid)

			cmds = append(cmds, paneResizeCmd(p.Grid[r][c], paneWidth, paneHeight))
		}
	}
	return cmds

}
