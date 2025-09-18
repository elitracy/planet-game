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

func (p Dashboard) GetId() int       { return p.id }
func (p *Dashboard) SetId(id int)    { p.id = id }
func (p Dashboard) GetTitle() string { return p.title }

func (p *Dashboard) Init() tea.Cmd { return nil }

func (p *Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return p, pushFocusCmd(p.Grid[p.ActiveRow][p.ActiveCol])
		case "esc":
			return p, popFocusCmd()
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

	activeStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("212")).
		Padding(1, 2)

	inactiveStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	render := make([][]string, len(p.Grid))
	for r := range p.Grid {
		render[r] = make([]string, len(p.Grid[r]))
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

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func NewDashboard(grid [][]int, activeRow, activeCol int, title string) *Dashboard {
	return &Dashboard{
		Grid:      grid,
		ActiveRow: activeRow,
		ActiveCol: activeCol,
		title:     title,
	}
}
