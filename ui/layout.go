package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dashboard struct {
	id        int
	title     string
	Grid      [][]tea.Model
	ActiveRow int
	ActiveCol int
}

var FocusStack []tea.Model

func PushFocus(pane tea.Model) {
	FocusStack = append(FocusStack, pane)
}

func PopFocus() tea.Model {
	if len(FocusStack) > 1 {
		FocusStack = FocusStack[:len(FocusStack)-1]
	}

	pane := FocusStack[len(FocusStack)-1]
	return pane
}

func ActivePane() tea.Model {
	if len(FocusStack) < 2 {
		return FocusStack[0]
	}

	return FocusStack[len(FocusStack)-1]
}

func (m Dashboard) GetId() int {
	return m.id
}

func (m Dashboard) GetTitle() string {
	return m.title
}

func (m *Dashboard) Init() tea.Cmd {
	var cmds []tea.Cmd

	for r := range m.Grid {
		for c := range m.Grid[r] {
			cmds = append(cmds, m.Grid[r][c].Init())
		}
	}

	return tea.Batch(cmds...)
}

func (m *Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			PushFocus(m.Grid[m.ActiveRow][m.ActiveCol])
		case "esc":
			PopFocus()
			return m, nil
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

	for r := range m.Grid {
		for c := range m.Grid[r] {
			var cmd tea.Cmd
			if r == m.ActiveRow && c == m.ActiveCol {
				m.Grid[r][c], cmd = m.Grid[r][c].Update(msg)
			} else {
				// send specific messages for background tasks
				switch msg := msg.(type) {
				case TickMsg:
					m.Grid[r][c], cmd = m.Grid[r][c].Update(msg)
				}
			}

			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}

	active := ActivePane()
	return active, tea.Batch(cmds...)
}

func (m *Dashboard) View() string {

	activeStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("212")).
		Padding(1, 2)

	inactiveStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	render := make([][]string, len(m.Grid))
	for r := range m.Grid {
		render[r] = make([]string, len(m.Grid[r]))
		for c := range m.Grid[r] {
			if r == m.ActiveRow && c == m.ActiveCol {
				render[r][c] = activeStyle.Render(m.Grid[r][c].View())
			} else {
				render[r][c] = inactiveStyle.Render(m.Grid[r][c].View())
			}
		}
	}

	var rows []string
	for r := range render {
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, render[r]...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func NewDashboard(grid [][]tea.Model, activeRow, activeCol, id int, title string) Dashboard {
	return Dashboard{
		Grid:      grid,
		ActiveRow: activeRow,
		ActiveCol: activeCol,
		id:        id,
		title:     title,
	}
}
