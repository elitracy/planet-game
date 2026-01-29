package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/orders"
	"github.com/elitracy/planets/state"
)

var orderStatusTypes = []events.EventStatus{
	events.EventPending,
	events.EventExecuting,
	events.EventComplete,
	events.EventFailed,
}

var orderStatuses [][]*orders.Order // eventStatus is the index

type OrderStatusListPane struct {
	*Pane
	cursor         int
	orderScheduler *events.EventScheduler[*orders.Order]
	theme          UITheme
}

func NewOrderStatusListPane(title string, orderScheduler *events.EventScheduler[*orders.Order]) *OrderStatusListPane {
	pane := &OrderStatusListPane{
		Pane: &Pane{
			title: title,
			keys:  NewKeyBindings(),
		},
		cursor:         0,
		orderScheduler: orderScheduler,
	}

	return pane
}

func (p *OrderStatusListPane) Init() tea.Cmd {
	p.keys.
		Set(Quit, "q").
		Set(Back, "esc").
		Set(Up, "k").
		Set(Down, "j")

	p.updateOrderStatusMap()

	orderList := NewOrderListPane(orderStatuses[p.cursor], events.EventStatus(p.cursor))
	paneID := PaneManager.AddPane(orderList)
	return tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
}

func (p *OrderStatusListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	p.updateOrderStatusMap()

	switch msg := msg.(type) {
	case paneResizeMsg:
		p.width = msg.width
		p.height = msg.height
	case tea.KeyMsg:
		switch msg.String() {
		case p.keys.Get(Up):
			if p.cursor > 0 {
				p.cursor--
			}

			logging.Info("status %v", events.EventStatus(p.cursor))
			orderList := NewOrderListPane(orderStatuses[p.cursor], events.EventStatus(p.cursor))
			paneID := PaneManager.AddPane(orderList)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case p.keys.Get(Down):
			if p.cursor < len(orderStatuses)-1 {
				p.cursor++
			}

			logging.Info("status %v", events.EventStatus(p.cursor))
			orderList := NewOrderListPane(orderStatuses[p.cursor], events.EventStatus(p.cursor))
			paneID := PaneManager.AddPane(orderList)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case p.keys.Get(Back):
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case p.keys.Get(Quit):
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *OrderStatusListPane) View() string {
	p.theme = GetPaneTheme(p)

	var rows []string

	rowIdx := 0
	for status := range orderStatuses {

		row := events.EventStatus(status).String()
		switch events.EventStatus(status) {
		case events.EventPending:
			if len(orderStatuses[status]) != 0 {
				row += "*"
			}
		case events.EventExecuting:
			if len(orderStatuses[status]) != 0 {
				row += "*"
			}
		}

		if p.cursor == rowIdx {
			row = p.theme.FocusedStyle.Render(row)
		} else {
			row = p.theme.BlurredStyle.Render(row)
		}
		rows = append(rows, row)

		rowIdx++
	}

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, rows...)

	title := p.Pane.Title()
	titleStyled := Style.Width(p.width).Bold(true).AlignHorizontal(lipgloss.Center).PaddingBottom(1).Render(title)

	content := lipgloss.JoinVertical(lipgloss.Left, titleStyled, infoContainer)

	return content
}

func (p *OrderStatusListPane) updateOrderStatusMap() {
	orderStatuses = make([][]*orders.Order, len(orderStatusTypes))

	for _, order := range p.orderScheduler.PriorityQueue {
		switch order.Status {
		case events.EventPending:
			orderStatuses[events.EventPending] = append(orderStatuses[events.EventPending], order)
		case events.EventExecuting:
			orderStatuses[events.EventExecuting] = append(orderStatuses[events.EventExecuting], order)
		case events.EventFailed:
			orderStatuses[events.EventFailed] = append(orderStatuses[events.EventFailed], order)
		}
	}

	for _, order := range state.State.CompletedOrders {
		switch order.Status {
		case events.EventComplete:
			orderStatuses[events.EventComplete] = append(orderStatuses[events.EventComplete], order)
		case events.EventFailed:
			orderStatuses[events.EventFailed] = append(orderStatuses[events.EventFailed], order)
		}
	}

}
