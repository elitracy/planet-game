package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/orders"
)

type OrderDetailsPane struct {
	*engine.Pane
	theme          UITheme
	orderInfoTable engine.ManagedPane
	order          *orders.Order
	progressBars   map[engine.EventID]engine.PaneID
}

func NewOrderDetailsPane(order *orders.Order) *OrderDetailsPane {
	return &OrderDetailsPane{
		Pane:  engine.NewPane(order.GetName(), engine.NewKeyBindings()),
		order: order,
	}
}

func (p *OrderDetailsPane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Back, "esc").
		Set(engine.Up, "k").
		Set(engine.Down, "j")

	keymaps := make(map[string]func() tea.Cmd)
	keymaps[p.GetKeys().Get(engine.Select)] = func() tea.Cmd {
		return tea.Sequence(pushDetailStackCmd(p.orderInfoTable.ID()), pushFocusStackCmd(p.orderInfoTable.ID()))
	}
	keymaps[p.GetKeys().Get(engine.Back)] = func() tea.Cmd {
		return tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
	}

	p.initProgressBars()

	infoTable := p.createInfoTable()
	p.orderInfoTable = NewInfoTablePane(
		infoTable,
		keymaps,
	)

	PaneManager.AddPane(p.orderInfoTable)
	for _, action := range p.order.Actions {
		engine.Info("actionID: %v", action.GetID())
	}
	return nil
}

func (p *OrderDetailsPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case paneResizeMsg:
		p.SetSize(msg.width, msg.height)

		msg.width = 15
		for _, paneID := range p.progressBars {
			progressBar := PaneManager.Panes[paneID]
			model, cmd := progressBar.Update(msg)
			cmds = append(cmds, cmd)
			progressBar = model.(ManagedPane)
		}
		return p, tea.Batch(cmds...)
	case engine.TickMsg:
		p.orderInfoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case game.UITickMsg:
		p.initProgressBars()
		p.orderInfoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	for _, paneID := range p.progressBars {
		progressBar := PaneManager.Panes[paneID]
		model, cmd := progressBar.Update(msg)
		cmds = append(cmds, cmd)
		progressBar = model.(ManagedPane)
	}

	model, cmd := p.orderInfoTable.Update(msg)
	cmds = append(cmds, cmd)
	p.orderInfoTable = model.(engine.ManagedPane)
	return p, tea.Batch(cmds...)
}

func (p *OrderDetailsPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.Title()
	titleStyled := Style.Width(p.Width()).AlignHorizontal(lipgloss.Center).Bold(true).PaddingBottom(1).Render(title)

	return lipgloss.JoinVertical(lipgloss.Left, titleStyled, p.orderInfoTable.View())
}

func (p OrderDetailsPane) createInfoTable() table.Model {
	infoTable := table.New(
		table.WithColumns(p.createColumns()),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(40),
	)

	return infoTable
}

func (p *OrderDetailsPane) createColumns() []table.Column {

	columns := []table.Column{
		{Title: "Action", Width: 30},
		{Title: "Start Time", Width: 15},
		{Title: "Completion Time", Width: 15},
		{Title: "Status", Width: 15},
	}
	if p.order.Status == engine.EventExecuting {
		columns = append(columns, table.Column{Title: "Progress", Width: 15})
	}

	return columns
}

func (p *OrderDetailsPane) createRows() []table.Row {

	rows := []table.Row{}
	for _, action := range p.order.GetActions() {
		start := action.GetStartTick().String()
		end := (action.GetStartTick() + action.GetDuration()).String()

		var row table.Row

		switch action.GetStatus() {
		case engine.EventPending:
			row = table.Row{action.GetDescription(), start, end, "Pending"}
		case engine.EventExecuting:
			paneID := p.progressBars[action.GetID()]
			row = table.Row{action.GetDescription(), start, end, "In Progress", PaneManager.Panes[paneID].View()}
		case engine.EventComplete:
			row = table.Row{action.GetDescription(), start, end, "Complete"}
		}

		rows = append(rows, row)

	}

	return rows
}

func (p *OrderDetailsPane) initProgressBars() {
	if p.progressBars == nil {
		p.progressBars = make(map[engine.EventID]engine.PaneID)
	}

	for _, action := range p.order.Actions {
		if _, ok := p.progressBars[action.GetID()]; !ok {
			progressBar := NewProgressBarPane(action.GetStartTick(), action.GetStartTick()+action.GetDuration())
			id := PaneManager.AddPane(progressBar)
			p.progressBars[action.GetID()] = id
			engine.Info("creating progress bar: %v", action.GetDescription())

		}
	}
}
