package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/models"
	"github.com/elitracy/planets/game/orders"
)

type StarSystemDetailsPane struct {
	*engine.Pane

	system          *models.StarSystem
	theme           UITheme
	systemInfoTable engine.ManagedPane
}

func NewSystemInfoPane(title string, system *models.StarSystem) *StarSystemDetailsPane {
	return &StarSystemDetailsPane{
		Pane:   engine.NewPane(title, engine.NewKeyBindings()),
		system: system,
	}
}

func (p *StarSystemDetailsPane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Back, "esc").
		Set(engine.Up, "k").
		Set(engine.Down, "j").
		Set(engine.Select, "enter")

	keymaps := make(map[string]func() tea.Cmd)

	keymaps[p.GetKeys().Get(engine.Back)] = func() tea.Cmd {
		return tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
	}
	keymaps[p.GetKeys().Get(engine.Select)] = func() tea.Cmd {
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
	case engine.TickMsg:
		p.systemInfoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case config.UITickMsg:
		p.systemInfoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case tea.KeyMsg:
		switch msg.String() {
		case p.GetKeys().Get(engine.Colonize):
			return p.handleColonizeOrder()
		case p.GetKeys().Get(engine.Scout):
			return p.handleScoutOrder()
		case "esc":
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	}
	model, cmd := p.systemInfoTable.Update(msg)
	cmds = append(cmds, cmd)
	p.systemInfoTable = model.(engine.ManagedPane)
	return p, tea.Batch(cmds...)
}

func (p *StarSystemDetailsPane) View() string {
	p.theme = GetPaneTheme(p)

	distance := engine.EuclidianDistance(game.State.Player.Position, p.system.Location.Position)
	distanceStyled := fmt.Sprintf(" (%v AU)", humanize.Comma(int64(distance)))
	distanceStyled = p.theme.DimmedStyle.Render(distanceStyled)

	title := p.system.GetName()
	titleStyled := Style.Bold(true).Render(title)

	header := lipgloss.JoinHorizontal(lipgloss.Top, titleStyled, distanceStyled)
	headerStyled := Style.Width(p.Width()).Align(lipgloss.Center).Bold(true).PaddingBottom(1).Render(header)

	p.GetKeys().Set(engine.Back, "esc")

	if !p.system.Planets[p.createInfoTable().Cursor()].Colonized {
		p.GetKeys().Set(engine.Colonize, "c")
	} else {
		p.GetKeys().Set(engine.Select, "enter")
	}

	if !p.system.Scouted {
		p.GetKeys().Set(engine.Scout, "s")
	} else {
		p.GetKeys().Set(engine.Colonize, "c")
	}

	if !p.system.Scouted && !p.system.Colonized {
		noDataMsg := Style.Width(p.Width()).AlignHorizontal(lipgloss.Center).Bold(true).Render("<No data for system>")
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

		radialDistance := engine.EuclidianDistance(planet.Location.Position, p.system.Location.Position)

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
		&game.State.ShipManager,
		func(ship *models.Ship) {
			order := orders.NewScoutDestinationOrder(
				ship,
				models.Location{
					Position: p.system.Location.Position,
					Entity:   p.system.Planets[p.createInfoTable().Cursor()],
				},
				game.State.CurrentTick+100)
			game.State.PushOrder(order)
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
