package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/logging"
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
		Padding(1, 2).
		Background(lipgloss.Color("235"))

	inactiveStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2).
		Background(lipgloss.Color("235"))

	prePaddingRender := make([][]string, len(m.Grid))
	for r := range m.Grid {
		prePaddingRender[r] = make([]string, len(m.Grid[r]))
		for c := range m.Grid[r] {
			content := m.Grid[r][c].View()
			if r == m.ActiveRow && c == m.ActiveCol {
				prePaddingRender[r][c] = activeStyle.Render(content)
			} else {
				prePaddingRender[r][c] = inactiveStyle.Render(content)
			}
		}
	}

	maxRowWidth := 0
	for r := range prePaddingRender {
		rowWidth := 0
		for c := range prePaddingRender[r] {
			w := lipgloss.Width(prePaddingRender[r][c])
			rowWidth += w
		}
		maxRowWidth = max(maxRowWidth, rowWidth)
	}

	// adjust padding for each cell to fit with max row width
	rendered := make([][]string, len(m.Grid))
	for r := range len(m.Grid) {
		rendered[r] = make([]string, len(m.Grid[r]))
		for c := range len(m.Grid[r]) {
			content := m.Grid[r][c].View()

			w := lipgloss.Width(prePaddingRender[r][c])

			rowPadding := maxRowWidth - len(m.Grid[r])
			paddingPerPane := rowPadding / w / 2

			logging.Log(fmt.Sprintf("Row: %d", r), "LAYOUT")
			logging.Log(fmt.Sprintf("Pre Padding Width: %d ", w), "LAYOUT")

			if r == m.ActiveRow && c == m.ActiveCol {
				rendered[r][c] = activeStyle.Copy().
					PaddingLeft(paddingPerPane).
					PaddingRight(paddingPerPane).
					Render(content)
			} else {
				rendered[r][c] = inactiveStyle.Copy().
					PaddingLeft(paddingPerPane).
					PaddingRight(paddingPerPane).
					Render(content)
			}
			logging.Log(fmt.Sprintf("Padding Width: %d ", lipgloss.Width(rendered[r][c])), "LAYOUT")

		}
	}

	var rows []string
	for r := range rendered {
		for c := range rendered[r] {
			w := lipgloss.Width(rendered[r][c])
			h := lipgloss.Height(rendered[r][c])

			rendered[r][c] = lipgloss.Place(
				w,
				h,
				lipgloss.Left,
				lipgloss.Top,
				rendered[r][c],
			)

		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rendered[r]...))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
