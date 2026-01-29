package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/orders"
	"github.com/elitracy/planets/state"
)

var loadingBarWidth = 0.5

type OrderListPane struct {
	*Pane
	cursor         int
	orders         []*orders.Order
	orderInfoTable ManagedPane
	progressBars   map[events.EventID]core.PaneID
	status         events.EventStatus
	theme          UITheme
}

func NewOrderListPane(orders []*orders.Order, status events.EventStatus) *OrderListPane {
	pane := &OrderListPane{
		Pane: &Pane{
			keys: NewKeyBindings(),
		},
		status: status,
		orders: orders,
	}

	return pane
}

func (p *OrderListPane) Init() tea.Cmd {
	p.keys.
		Set(Quit, "q").
		Set(Back, "esc").
		Set(Up, "k").
		Set(Down, "j")

	p.initProgrssBars()

	keymaps := make(map[string]func() tea.Cmd)
	keymaps[p.keys.Get(Back)] = func() tea.Cmd {
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

	switch msg := msg.(type) {
	case paneResizeMsg:
		p.width = msg.width
		p.height = msg.height

		msg.width = 25
		for _, paneID := range p.progressBars {
			progressBar := PaneManager.Panes[paneID]
			model, cmd := progressBar.Update(msg)
			cmds = append(cmds, cmd)
			progressBar = model.(ManagedPane)
		}
		return p, tea.Batch(cmds...)
	case core.TickMsg:
		p.orderInfoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case core.UITickMsg:
		p.orderInfoTable.(*InfoTablePane).table.SetRows(p.createRows())
		p.orderInfoTable.(*InfoTablePane).table.SetColumns(p.createColumns())
	case tea.KeyMsg:
		switch msg.String() {
		case p.keys.Get(Up):
			if p.cursor > 0 {
				p.cursor--
			}
		case p.keys.Get(Down):
			if p.cursor < len(p.orders) {
				p.cursor++
			}
		case p.keys.Get(Back):
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case p.keys.Get(Quit):
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
	p.orderInfoTable = model.(ManagedPane)

	return p, tea.Batch(cmds...)
}

func (p *OrderListPane) View() string {
	p.theme = GetPaneTheme(p)

	title := fmt.Sprintf("%v Orders", p.status)
	titleStyled := Style.Width(p.width).Bold(true).Align(lipgloss.Center).PaddingBottom(1).Render(title)

	content := lipgloss.JoinVertical(lipgloss.Left, titleStyled, p.orderInfoTable.View())

	return content
}

func (p *OrderListPane) initProgrssBars() {
	if p.progressBars == nil {
		p.progressBars = make(map[events.EventID]core.PaneID)
	}

	for _, order := range p.orders {
		if _, ok := p.progressBars[order.ID]; !ok {
			progressBar := NewLoadingBarPane(order.GetExecuteTick(), order.GetEndTick())
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
		table.WithHeight(10),
	)

	return infoTable
}

func (p *OrderListPane) createColumns() []table.Column {

	switch p.status {
	case events.EventPending:
		return []table.Column{
			{Title: "Order", Width: 15},
			{Title: "Time to Execution", Width: 25},
		}
	case events.EventExecuting:
		return []table.Column{
			{Title: "Order", Width: 15},
			{Title: "Completion Time", Width: 25},
			{Title: "Progress", Width: p.width - 40},
		}
	default:
		return []table.Column{
			{Title: "Order", Width: 15},
		}
	}

}

func (p *OrderListPane) createRows() []table.Row {

	rows := []table.Row{}
	for _, order := range p.orders {
		if order.GetStatus() != p.status {
			continue
		}

		switch p.status {
		case events.EventPending:
			duration := (order.GetExecuteTick() - state.State.CurrentTick).ToDuration(consts.TICKS_PER_SECOND)

			row := table.Row{order.GetName(), humanize.Time(time.Now().Add(duration))}
			rows = append(rows, row)
		case events.EventExecuting:
			progressBar := PaneManager.Panes[p.progressBars[order.GetID()]]
			duration := (order.GetEndTick() - state.State.CurrentTick).ToDuration(consts.TICKS_PER_SECOND)

			row := table.Row{order.GetName(), humanize.Time(time.Now().Add(duration)), progressBar.View()}
			rows = append(rows, row)
		default:
			row := table.Row{order.GetName()}
			rows = append(rows, row)
		}

	}

	return rows
}
