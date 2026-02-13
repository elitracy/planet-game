package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/orders"
)

type OrderListPane struct {
	*engine.Pane
	cursor         int
	filteredOrders []*orders.Order
	orderInfoTable engine.ManagedPane
	progressBars   map[engine.EventID]engine.PaneID
	status         engine.EventStatus
	theme          UITheme
}

func NewOrderListPane(status engine.EventStatus) *OrderListPane {
	pane := &OrderListPane{
		Pane:   engine.NewPane("Order List", engine.NewKeyBindings()),
		status: status,
	}

	return pane
}

func (p *OrderListPane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Back, "esc").
		Set(engine.Select, "enter").
		Set(engine.Up, "k").
		Set(engine.Down, "j")

	p.filterOrders()
	p.initProgressBars()

	keymaps := make(map[string]func() tea.Cmd)
	keymaps[p.GetKeys().Get(engine.Select)] = func() tea.Cmd {
		order := p.filteredOrders[p.orderInfoTable.(*InfoTablePane).table.Cursor()]
		pane := NewOrderDetailsPane(order)
		paneID := PaneManager.AddPane(pane)

		return tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
	}
	keymaps[p.GetKeys().Get(engine.Back)] = func() tea.Cmd {
		return tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
	}

	infoTable := p.createInfoTable()
	p.orderInfoTable = NewInfoTablePane(
		infoTable,
		keymaps,
	)

	PaneManager.AddPane(p.orderInfoTable)
	return nil
}

func (p *OrderListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	if PaneManager.PeekFocusStack() == p.ID() && len(p.filteredOrders) == 0 {
		cmds = append(cmds, popFocusStackCmd())
	}

	switch msg := msg.(type) {
	case paneResizeMsg:
		p.SetSize(msg.width, msg.height)

		msg.width = 30
		for _, paneID := range p.progressBars {
			progressBar := PaneManager.Panes[paneID]
			model, cmd := progressBar.Update(msg)
			cmds = append(cmds, cmd)
			progressBar = model.(engine.ManagedPane)
		}
		return p, tea.Batch(cmds...)
	case engine.TickMsg:
		p.filterOrders()
		p.initProgressBars()

		if len(p.filteredOrders) == 0 {
			p.GetKeys().Unset(engine.Select)
		} else {
			p.GetKeys().Set(engine.Select, "enter")
		}

		p.orderInfoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
		p.orderInfoTable.(*InfoTablePane).table.SetRows(p.createRows())
		p.orderInfoTable.(*InfoTablePane).table.SetColumns(p.createColumns())
	case tea.KeyMsg:
		switch msg.String() {
		case p.GetKeys().Get(engine.Up):
			if p.cursor > 0 {
				p.cursor--
			}
		case p.GetKeys().Get(engine.Down):
			if p.cursor < len(p.filteredOrders) {
				p.cursor++
			}
		case p.GetKeys().Get(engine.Back):
			return p, tea.Sequence(popFocusStackCmd())
		case p.GetKeys().Get(engine.Quit):
			return p, tea.Quit
		}
	}

	for _, paneID := range p.progressBars {
		progressBar := PaneManager.Panes[paneID]
		model, cmd := progressBar.Update(msg)
		cmds = append(cmds, cmd)
		progressBar = model.(engine.ManagedPane)
	}

	model, cmd := p.orderInfoTable.Update(msg)
	cmds = append(cmds, cmd)
	p.orderInfoTable = model.(engine.ManagedPane)

	return p, tea.Batch(cmds...)
}

func (p *OrderListPane) View() string {
	p.theme = GetPaneTheme(p)

	title := fmt.Sprintf("%v Orders", p.status)
	titleStyled := Style.Width(p.Width()).Bold(true).Align(lipgloss.Center).PaddingBottom(1).Render(title)

	content := lipgloss.JoinVertical(lipgloss.Left, titleStyled, p.orderInfoTable.View())

	return content
}

func (p *OrderListPane) initProgressBars() {
	if p.progressBars == nil {
		p.progressBars = make(map[engine.EventID]engine.PaneID)
	}

	for _, order := range p.filteredOrders {
		if order.GetStatus() != p.status {
			continue
		}
		if _, ok := p.progressBars[order.ID]; !ok {
			progressBar := NewProgressBarPane(order.GetStartTick(), order.GetStartTick()+order.GetDuration())
			id := PaneManager.AddPane(progressBar)
			p.progressBars[order.GetID()] = id

		}
	}
}

func (p OrderListPane) createInfoTable() table.Model {
	infoTable := table.New(
		table.WithColumns(p.createColumns()),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(40),
	)

	return infoTable
}

func (p *OrderListPane) createColumns() []table.Column {

	switch p.status {
	case engine.EventPending:
		return []table.Column{
			{Title: "Order", Width: 25},
			{Title: "Time to Execution", Width: 40},
		}
	case engine.EventExecuting:
		return []table.Column{
			{Title: "Order", Width: 25},
			{Title: "Completion Time", Width: 25},
			{Title: "Progress", Width: 40},
		}
	default:
		return []table.Column{
			{Title: "Order", Width: 25},
		}
	}

}

func (p *OrderListPane) createRows() []table.Row {

	rows := []table.Row{}
	for _, order := range p.filteredOrders {
		if order.GetStatus() != p.status {
			continue
		}

		switch p.status {
		case engine.EventPending:
			duration := (order.GetStartTick() - game.State.CurrentTick).ToDuration(engine.TICKS_PER_SECOND)

			row := table.Row{order.GetName(), humanize.Time(time.Now().Add(duration))}
			rows = append(rows, row)
		case engine.EventExecuting:
			progressBar := PaneManager.Panes[p.progressBars[order.GetID()]]
			duration := (order.GetStartTick() + order.GetDuration() - game.State.CurrentTick).ToDuration(engine.TICKS_PER_SECOND)

			row := table.Row{order.GetName(), humanize.Time(time.Now().Add(duration)), progressBar.View()}
			rows = append(rows, row)
		default:
			row := table.Row{order.GetName()}
			rows = append(rows, row)
		}
	}

	return rows
}

func (p *OrderListPane) filterOrders() {
	p.filteredOrders = []*orders.Order{}

	if p.status == engine.EventComplete {
		p.filteredOrders = game.State.CompletedOrders
		return
	}

	for _, order := range game.State.OrderScheduler.PriorityQueue {
		if order.GetStatus() == p.status {
			p.filteredOrders = append(p.filteredOrders, order)
		}
	}

}
