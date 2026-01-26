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

const (
	DEFAULT_KEYS = "Back: esc"
	INFO_KEY     = " | Info: enter"
	SCOUT_KEY    = " | Scout: s"
	COLONIZE_KEY = " | Colonize: c"
)

type StarSystemInfoPane struct {
	*Pane

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
		planet := p.system.Planets[cursor]

		if !planet.Scouted && !planet.Colonized {
			return nil
		}

		pane := NewPlanetInfoPane("Planet Info", p.system.Planets[cursor])
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
		case "s":
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

func (p *StarSystemInfoPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.system.Name
	titleStyled := Style.Width(p.width).Align(lipgloss.Center).Bold(true).PaddingBottom(1).Render(title)
	p.keys = DEFAULT_KEYS

	if !p.system.Planets[p.createInfoTable().Cursor()].Colonized {
		p.keys += COLONIZE_KEY
	} else {
		p.keys += INFO_KEY
	}

	if !p.system.Scouted {
		p.keys += SCOUT_KEY
	} else {
		p.keys += INFO_KEY
	}

	return lipgloss.JoinVertical(lipgloss.Left, titleStyled, p.systemInfoTable.View())
}

func NewSystemInfoPane(title string, system *models.StarSystem) *StarSystemInfoPane {
	return &StarSystemInfoPane{
		Pane: &Pane{
			title: title,
			keys:  "Planet Info: enter | Back: esc | ",
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

		var row table.Row
		if planet.Scouted || planet.Colonized {
			row = table.Row{planet.Name, planet.Position.String(), populationString}
		} else {
			row = table.Row{planet.Name, planet.Position.String(), ""}
		}

		rows = append(rows, row)
	}

	return rows
}

func (p *StarSystemInfoPane) handleScoutOrder() (tea.Model, tea.Cmd) {
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
