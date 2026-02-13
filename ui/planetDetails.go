package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/models"
)

type PlanetDetailsPane struct {
	*engine.Pane

	infoTable engine.ManagedPane
	planet    *models.Planet
	theme     UITheme
}

func (p *PlanetDetailsPane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Back, "esc").
		Set(engine.Up, "k").
		Set(engine.Down, "j").
		Set(engine.Colonize, "c").
		Set(engine.Quit, "q")

	keymaps := make(map[string]func() tea.Cmd)

	keymaps[p.GetKeys().Get(engine.Back)] = func() tea.Cmd {
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
func (p *PlanetDetailsPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case engine.TickMsg:
		p.infoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case config.UITickMsg:
		p.infoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case paneResizeMsg:
		p.SetSize(msg.width, msg.height)
	case tea.KeyMsg:
		switch msg.String() {
		case p.GetKeys().Get(engine.Colonize):
			return p.handleColonization()
		case p.GetKeys().Get(engine.Back):
			return p, tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
		case p.GetKeys().Get(engine.Quit):
			return p, tea.Quit
		}
	}

	model, cmd := p.infoTable.Update(msg)
	p.infoTable = model.(engine.ManagedPane)

	cmds = append(cmds, cmd)

	return p, tea.Batch(cmds...)
}

func (p *PlanetDetailsPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.planet.Name

	population := fmt.Sprintf("Population: %v", humanize.Comma(int64(p.planet.Population)))
	populationStyled := p.theme.DimmedStyle.Width(p.Width()).AlignHorizontal(lipgloss.Center).Render(population)

	titleStyled := Style.Width(p.Width()).Align(lipgloss.Center).Bold(true).Render(title)

	header := lipgloss.JoinVertical(lipgloss.Center, titleStyled, populationStyled)
	headerStyled := Style.PaddingBottom(1).Render(header)

	content := lipgloss.JoinVertical(lipgloss.Left, headerStyled, p.infoTable.View())

	return content

}

func NewPlanetDetailsPane(title string, planet *models.Planet) *PlanetDetailsPane {

	return &PlanetDetailsPane{
		Pane:   engine.NewPane(title, engine.NewKeyBindings()),
		planet: planet,
	}
}

func (p *PlanetDetailsPane) createInfoTable() table.Model {

	columns := []table.Column{
		{Title: "Stat", Width: 15},
		{Title: "Quantity", Width: 10},
		{Title: "Rate (units/p)", Width: 15},
	}

	infoTable := table.New(
		table.WithColumns(columns),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(12),
	)

	return infoTable
}

func (p *PlanetDetailsPane) createRows() []table.Row {

	rows := []table.Row{
		{"Food", humanize.Comma(int64(p.planet.Food.Quantity)), humanize.Comma(int64(p.planet.Food.ConsumptionRate))},
		{"Mineral", humanize.Comma(int64(p.planet.Minerals.Quantity)), humanize.Comma(int64(p.planet.Minerals.ConsumptionRate))},
		{"Energy", humanize.Comma(int64(p.planet.Energy.Quantity)), humanize.Comma(int64(p.planet.Energy.ConsumptionRate))},

		{"", "", ""},
		{"Farms", humanize.Comma(int64(len(p.planet.Farms))), humanize.Comma(int64(p.planet.GetFarmProduction()))},
		{"Mines", humanize.Comma(int64(len(p.planet.Mines))), humanize.Comma(int64(p.planet.GetMineProduction()))},
		{"Solar Grids", humanize.Comma(int64(len(p.planet.SolarGrids))), humanize.Comma(int64(p.planet.GetSolarGridProduction()))},

		{"", "", ""},
		{"Happiness", fmt.Sprintf("%v%%", int64(p.planet.Happiness.Quantity*100)), strconv.FormatFloat(float64(p.planet.Happiness.GrowthRate), 'f', 2, 32)},
		{"Corruption", fmt.Sprintf("%v%%", int64(p.planet.Corruption.Quantity*100)), strconv.FormatFloat(float64(p.planet.Corruption.GrowthRate), 'f', 2, 32)},
		{"Unrest", fmt.Sprintf("%v%%", int64(p.planet.Unrest.Quantity*100)), strconv.FormatFloat(float64(p.planet.Unrest.GrowthRate), 'f', 2, 32)},
	}

	return rows
}

func (p *PlanetDetailsPane) handleColonization() (tea.Model, tea.Cmd) {
	pane := NewCreateColonyPane(
		"Order Colonization: "+p.planet.Name,
		p.planet,
	)

	paneID := PaneManager.AddPane(pane)
	return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
}
