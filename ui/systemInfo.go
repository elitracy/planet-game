package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
)

type StarSystemInfoPane struct {
	*Pane

	cursor          int
	system          *models.StarSystem
	theme           UITheme
	systemInfoTable ManagedPane
}

func (p *StarSystemInfoPane) Init() tea.Cmd {

	keymaps := make(map[string]func() tea.Cmd)

	keymaps["esc"] = func() tea.Cmd {
		return tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
	}
	keymaps["enter"] = func() tea.Cmd {
		cursor := p.systemInfoTable.(*InfoTablePane).table.Cursor()

		pane := &PlanetInfoPane{
			Pane: &Pane{
				title: "Planet Info",
			},
			planet: p.system.Planets[cursor],
		}
		paneID := PaneManager.AddPane(pane)
		return tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
	}

	infoTable := p.createInfoTable()
	p.systemInfoTable = NewInfoTablePane(
		infoTable,
		keymaps,
	)

	PaneManager.AddPane(p.systemInfoTable)

	return nil
}
func (p *StarSystemInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case core.TickMsg:
		p.systemInfoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case core.UITickMsg:
		p.systemInfoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		case "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "j":
			if p.cursor < len(p.system.Planets)-1 {
				p.cursor++
			}
		}

	}
	model, cmd := p.systemInfoTable.Update(msg)
	cmds = append(cmds, cmd)
	p.systemInfoTable = model.(ManagedPane)
	return p, tea.Batch(cmds...)
}

func (p *StarSystemInfoPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.system.Name

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, title, p.systemInfoTable.View())

	scoutButton := "Scout : s"
	colonizeButton := "Colonize: c"

	buttons := lipgloss.JoinHorizontal(lipgloss.Left, scoutButton, " | ", colonizeButton)
	buttons = Style.Width(p.width).Border(lipgloss.NormalBorder(), true, false, false, false).Render(buttons)

	topContent := lipgloss.Place(
		p.width,
		p.height-lipgloss.Height(buttons),
		lipgloss.Left,
		lipgloss.Top,
		infoContainer,
	)

	content := lipgloss.JoinVertical(lipgloss.Left, topContent, buttons)
	return content

}

func NewSystemInfoPane(title string, system *models.StarSystem) *StarSystemInfoPane {
	return &StarSystemInfoPane{
		Pane: &Pane{
			title: title,
		},
		system: system,
	}
}

func (p StarSystemInfoPane) createInfoTable() table.Model {

	columns := []table.Column{
		{Title: "Planet", Width: 15},
		{Title: "Position (x,y,z)", Width: 20},
		{Title: "Population (Δpop/pulse)", Width: 25},
	}

	infoTable := table.New(
		table.WithColumns(columns),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(len(p.system.Planets)+1),
	)

	return infoTable
}

func (p *StarSystemInfoPane) createRows() []table.Row {

	rows := []table.Row{}
	for _, planet := range p.system.Planets {
		populationString := fmt.Sprintf("%v (%v)", humanize.Comma(int64(planet.Population)), strconv.Itoa(planet.PopulationGrowthRate))
		rows = append(rows, table.Row{planet.Name, planet.Position.String(), populationString})
	}

	return rows
}
