package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
)

const (
	BORDER_WIDTH = 2
)

type DashboardPane struct {
	*engine.Pane
	Grid      [][]engine.PaneID
	ActiveRow int
	ActiveCol int
}

func (p *DashboardPane) Init() tea.Cmd { return nil }

func (p *DashboardPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return p, setMainFocusCmd(p.Grid[p.ActiveRow][p.ActiveCol])
		case "esc":
			return p, popMainFocusCmd(p.Pane.ID())
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

func (p *DashboardPane) View() string {

	render := make([][]string, len(p.Grid))
	for r := range p.Grid {
		render[r] = make([]string, len(p.Grid[r]))

		for c := range p.Grid[r] {
			paneID := p.Grid[r][c]

			activeStyle := Style.
				Width(p.Pane.Width()).
				Height(p.Pane.Height() - BORDER_WIDTH).
				Border(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color("212"))

			inactiveStyle := Style.
				Width(p.Pane.Width()).
				Height(p.Pane.Height() - BORDER_WIDTH).
				Border(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color("240"))
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

func NewDashboard(grid [][]engine.PaneID, title string) *DashboardPane {
	return &DashboardPane{
		Grid: grid,
		Pane: engine.NewPane(title, nil),
	}
}
