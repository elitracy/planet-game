package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Pane struct {
	Id    int
	Label string
}

func (p Pane) Init() tea.Cmd                           { return nil }
func (p Pane) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }
func (p Pane) View() string                            { return p.Label }

type Dashboard struct {
	Grid      [][]Pane
	ActiveRow int
	ActiveCol int
}

func (m Dashboard) Init() tea.Cmd { return nil }

func (m Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			if m.ActiveCol > 0 {
				m.ActiveCol--
			}
		case "l":
			if m.ActiveCol < len(m.Grid[m.ActiveRow])-1 {
				m.ActiveCol++
			}
		case "k":
			if m.ActiveRow > 0 {
				if m.ActiveCol >= len(m.Grid[m.ActiveRow-1]) {
					m.ActiveCol = len(m.Grid[m.ActiveRow-1]) - 1
				}

				m.ActiveRow--
			}
		case "j":
			if m.ActiveRow < len(m.Grid)-1 {
				if m.ActiveCol >= len(m.Grid[m.ActiveRow+1]) {
					m.ActiveCol = len(m.Grid[m.ActiveRow+1]) - 1
				}
				m.ActiveRow++
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m Dashboard) View() string {

	activeStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("212")).
		Padding(1, 2)

	inactiveStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	rendered := make([][]string, len(m.Grid))
	for r := range m.Grid {
		rendered[r] = make([]string, len(m.Grid[r]))
		for c := range m.Grid[r] {
			content := m.Grid[r][c].View()
			if r == m.ActiveRow && c == m.ActiveCol {
				rendered[r][c] = activeStyle.Render(content)
			} else {
				rendered[r][c] = inactiveStyle.Render(content)
			}
		}
	}

	maxCols := 0
	for i := range len(m.Grid) {
		maxCols = max(len(m.Grid[i]), maxCols)
	}

	colWidths := make([]int, maxCols)
	rowHeights := make([]int, len(rendered))

	for r := range rendered {
		for c := range rendered[r] {
			w := lipgloss.Width(rendered[r][c])
			h := lipgloss.Height(rendered[r][c])

			if w > colWidths[c] {
				colWidths[c] = w
			}

			if h > rowHeights[r] {
				rowHeights[r] = h
			}
		}
	}

	for r := range rendered {
		for c := range rendered[r] {
			rendered[r][c] = lipgloss.Place(
				colWidths[c],
				rowHeights[r],
				lipgloss.Center,
				lipgloss.Center,
				rendered[r][c],
			)
		}
	}

	var rows []string
	for r := range rendered {
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rendered[r]...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
