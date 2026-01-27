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
	"github.com/elitracy/planets/models/events/orders"
	"github.com/elitracy/planets/state"
)

type StarSystemDetailsPane struct {
	*Pane

	system          *models.StarSystem
	theme           UITheme
	systemInfoTable ManagedPane
}

func NewSystemInfoPane(title string, system *models.StarSystem) *StarSystemDetailsPane {
	return &StarSystemDetailsPane{
		Pane: &Pane{
			title: title,
			keys:  NewKeyBindings(),
		},
		system: system,
	}
}

func (p *StarSystemDetailsPane) Init() tea.Cmd {
	p.keys.
		Set(Quit, "q").
		Set(Back, "esc").
		Set(Up, "k").
		Set(Down, "j").
		Set(Select, "enter")

	keymaps := make(map[string]func() tea.Cmd)

	keymaps[p.keys.Get(Back)] = func() tea.Cmd {
		return tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
	}
	keymaps[p.keys.Get(Select)] = func() tea.Cmd {
		cursor := p.systemInfoTable.(*InfoTablePane).table.Cursor()
		planet := p.system.Planets[cursor]

		if !planet.Scouted && !planet.Colonized {
			return nil
		}

		pane := NewPlanetDetailsPane("Planet Info", p.system.Planets[cursor])
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
func (p *StarSystemDetailsPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case core.TickMsg:
		p.systemInfoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case core.UITickMsg:
		p.systemInfoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case tea.KeyMsg:
		switch msg.String() {
		case p.keys.Get(Colonize):
			return p.handleColonizeOrder()
		case p.keys.Get(Scout):
			return p.handleScoutOrder()
		case "esc":
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	}
	model, cmd := p.systemInfoTable.Update(msg)
	cmds = append(cmds, cmd)
	p.systemInfoTable = model.(ManagedPane)
	return p, tea.Batch(cmds...)
}

func (p *StarSystemDetailsPane) View() string {
	p.theme = GetPaneTheme(p)

	distance := core.EuclidianDistance(state.State.Player.Position, p.system.Position)
	distanceStyled := fmt.Sprintf(" (%v AU)", humanize.Comma(int64(distance)))
	distanceStyled = p.theme.DimmedStyle.Render(distanceStyled)

	title := p.system.GetName()
	titleStyled := Style.Bold(true).Render(title)

	header := lipgloss.JoinHorizontal(lipgloss.Top, titleStyled, distanceStyled)
	headerStyled := Style.Width(p.width).Align(lipgloss.Center).Bold(true).PaddingBottom(1).Render(header)

	p.keys.Set(Back, "esc")

	if !p.system.Planets[p.createInfoTable().Cursor()].Colonized {
		p.keys.Set(Colonize, "c")
	} else {
		p.keys.Set(Select, "enter")
	}

	if !p.system.Scouted {
		p.keys.Set(Scout, "s")
	} else {
		p.keys.Set(Colonize, "c")
	}

	if !p.system.Scouted && !p.system.Colonized {
		noDataMsg := Style.Width(p.width).AlignHorizontal(lipgloss.Center).Bold(true).Render("<No Data for System>")
		return lipgloss.JoinVertical(lipgloss.Left, headerStyled, noDataMsg)
	}

	return lipgloss.JoinVertical(lipgloss.Left, headerStyled, p.systemInfoTable.View())
}

func (p StarSystemDetailsPane) createInfoTable() table.Model {
	infoTable := table.New(
		table.WithColumns(p.createColumns()),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(len(p.system.Planets)+1),
	)

	return infoTable
}
func (p *StarSystemDetailsPane) createColumns() []table.Column {

	return []table.Column{
		{Title: "Planet", Width: 15},
		{Title: "Orbit Distance (AU)", Width: 20},
		{Title: "Population (Δpop/pulse)", Width: 25},
	}
}

func (p *StarSystemDetailsPane) createRows() []table.Row {

	rows := []table.Row{}
	for _, planet := range p.system.Planets {
		populationString := fmt.Sprintf("%v (%v)", humanize.Comma(int64(planet.Population)), strconv.Itoa(planet.PopulationGrowthRate))

		radialDistance := core.EuclidianDistance(planet.Position, p.system.Position)

		var row table.Row
		if planet.Scouted || planet.Colonized {
			row = table.Row{planet.Name, humanize.Comma(int64(radialDistance)), populationString}
		} else {
			row = table.Row{planet.Name, humanize.Comma(int64(radialDistance)), ""}
		}

		rows = append(rows, row)
	}

	return rows
}

func (p *StarSystemDetailsPane) handleScoutOrder() (tea.Model, tea.Cmd) {
	pane := CreateNewShipManagementPane(
		"Ship Management",
		&state.State.ShipManager,
		func(ship *models.Ship) {
			order := orders.NewScoutDestinationOrder(ship, models.Destination{Position: p.system.Position, Entity: p.system.Planets[p.createInfoTable().Cursor()]}, state.State.CurrentTick+40)
			state.State.OrderScheduler.Push(order)
		},
	)

	paneID := PaneManager.AddPane(pane)
	return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
}

func (p *StarSystemDetailsPane) handleColonizeOrder() (tea.Model, tea.Cmd) {
	cursor := p.systemInfoTable.(*InfoTablePane).table.Cursor()
	pane := NewCreateColonyPane(
		"Order Colonization: "+p.system.Planets[cursor].Name,
		p.system.Planets[cursor],
	)

	paneID := PaneManager.AddPane(pane)
	return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
}
