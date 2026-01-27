package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
)

type PlanetInfoPane struct {
	*Pane

	infoTable ManagedPane
	planet    *models.Planet
	theme     UITheme
}

func (p *PlanetInfoPane) Init() tea.Cmd {
	p.keys.
		Set(Back, "esc").
		Set(Colonize, "c").
		Set(Quit, "q")

	keymaps := make(map[string]func() tea.Cmd)

	keymaps[p.keys.Get(Back)] = func() tea.Cmd {
		return tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
	}

	infoTable := p.createInfoTable()
	p.infoTable = NewInfoTablePane(
		infoTable,
		keymaps,
	)
	PaneManager.AddPane(p.infoTable)

	p.infoTable.Init()

	return nil
}
func (p *PlanetInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case core.UITickMsg:
		p.infoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case paneResizeMsg:
		p.height = msg.height
		p.width = msg.width
	case tea.KeyMsg:
		switch msg.String() {
		case p.keys.Get(Colonize):
			return p.handleColonization()
		case p.keys.Get(Back):
			return p, tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
		case p.keys.Get(Quit):
			return p, tea.Quit
		}
	}

	model, cmd := p.infoTable.Update(msg)
	p.infoTable = model.(ManagedPane)

	cmds = append(cmds, cmd)

	return p, tea.Batch(cmds...)
}

func (p *PlanetInfoPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.planet.Name

	population := fmt.Sprintf("Population: %d", p.planet.Population)

	titleStyled := Style.Width(p.width).Align(lipgloss.Center).Bold(true).PaddingBottom(1).Render(title)

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, population, p.infoTable.View())

	content := lipgloss.JoinVertical(lipgloss.Left, titleStyled, infoContainer)

	return content

}

func NewPlanetInfoPane(title string, planet *models.Planet) *PlanetInfoPane {

	return &PlanetInfoPane{
		Pane: &Pane{
			title: title,
			keys:  NewKeyBindings(),
		},
		planet: planet,
	}
}

func (p *PlanetInfoPane) createInfoTable() table.Model {

	columns := []table.Column{
		{Title: "Stat", Width: 15},
		{Title: "Quantity", Width: 10},
		{Title: "Rate", Width: 4},
	}

	infoTable := table.New(
		table.WithColumns(columns),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	return infoTable
}

func (p *PlanetInfoPane) createRows() []table.Row {

	rows := []table.Row{
		{"Food", strconv.Itoa(p.planet.Food.Quantity), strconv.Itoa(p.planet.Food.ConsumptionRate)},
		{"Minerals", strconv.Itoa(p.planet.Minerals.Quantity), strconv.Itoa(p.planet.Minerals.ConsumptionRate)},
		{"Energy", strconv.Itoa(p.planet.Energy.Quantity), strconv.Itoa(p.planet.Energy.ConsumptionRate)},

		{"", "", ""},
		{"Farms", strconv.Itoa(len(p.planet.Farms)), strconv.Itoa(p.planet.GetFarmProduction())},
		{"Mines", strconv.Itoa(len(p.planet.Mines)), strconv.Itoa(p.planet.GetMineProduction())},
		{"Solar Grids", strconv.Itoa(len(p.planet.SolarGrids)), strconv.Itoa(p.planet.GetSolarGridProduction())},

		{"", "", ""},
		{"Happiness", strconv.FormatFloat(float64(p.planet.Happiness.Quantity), 'f', 2, 32), strconv.FormatFloat(float64(p.planet.Happiness.GrowthRate), 'f', 2, 32)},
		{"Corruption", strconv.FormatFloat(float64(p.planet.Corruption.Quantity), 'f', 2, 32), strconv.FormatFloat(float64(p.planet.Corruption.GrowthRate), 'f', 2, 32)},
		{"Unrest", strconv.FormatFloat(float64(p.planet.Unrest.Quantity), 'f', 2, 32), strconv.FormatFloat(float64(p.planet.Unrest.GrowthRate), 'f', 2, 32)},
	}

	return rows
}

func (p *PlanetInfoPane) handleColonization() (tea.Model, tea.Cmd) {
	pane := NewCreateColonyPane(
		"Order Colonization: "+p.planet.Name,
		p.planet,
	)

	paneID := PaneManager.AddPane(pane)
	return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
}
