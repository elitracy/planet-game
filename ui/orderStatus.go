package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/orders"
	"github.com/elitracy/planets/state"
)

var (
	activeRowStyle   = Style.Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("212")).Padding(0, 1)
	inactiveRowStyle = Style.Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).Padding(0, 1)
	progressBarWidth int
)

type OrderStatusPane struct {
	*Pane
	cursor         int
	orderScheduler *events.EventScheduler[*orders.Order]
	progressBars   map[events.EventID]core.PaneID
}

func NewOrderStatusPane(title string, orderScheduler *events.EventScheduler[*orders.Order]) *OrderStatusPane {
	pane := &OrderStatusPane{
		Pane: &Pane{
			title: title,
		},
		orderScheduler: orderScheduler,
	}

	return pane
}

func (p *OrderStatusPane) Init() tea.Cmd {
	p.initProgrssBars()
	return nil
}

func (p *OrderStatusPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		var cmds []tea.Cmd
		for _, val := range p.progressBars {
			cmds = append(cmds, paneResizeCmd(val, progressBarWidth, msg.height))
		}
		return p, tea.Batch(cmds...)
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.orderScheduler.PriorityQueue)+len(state.State.CompletedOrders)-1 {
				p.cursor++
			}
		case "esc":
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	var cmds []tea.Cmd
	for _, paneID := range p.progressBars {
		progressBar := PaneManager.Panes[paneID]
		model, cmd := progressBar.Update(msg)
		cmds = append(cmds, cmd)
		progressBar = model.(ManagedPane)
	}

	return p, tea.Batch(cmds...)
}

func (p *OrderStatusPane) View() string {

	var pendingOrders []*orders.Order
	var executingOrders []*orders.Order
	var completedOrders []*orders.Order

	for _, order := range p.orderScheduler.PriorityQueue {
		switch order.GetStatus() {
		case consts.EventPending:
			pendingOrders = append(pendingOrders, order)
		case consts.EventExecuting:
			executingOrders = append(executingOrders, order)
		}
	}

	for _, order := range state.State.CompletedOrders {
		completedOrders = append(completedOrders, order)
	}

	// p.initProgrssBars()

	title := Style.Width(p.Pane.width).AlignHorizontal(lipgloss.Center).Render(p.Pane.title + "\n")

	var pendingOrderRows []string

	pendingOrdersTitle := "Pending"
	pendingOrdersTitleStyles := Style.Bold(true)
	pendingOrderRows = append(pendingOrderRows, pendingOrdersTitleStyles.Render(pendingOrdersTitle))

	currentOrder := 0
	for _, order := range pendingOrders {
		row := fmt.Sprintf("[%v] %v", order.GetStatus(), order.GetName())

		countDown := fmt.Sprintf("ETA: %vs", (order.GetExecuteTick()-state.State.CurrentTick)/core.TICKS_PER_SECOND)

		if lipgloss.Width(countDown)+lipgloss.Width(row) > p.Pane.width-5 {
			row = Style.Render(row)
			countDown = Theme.BlurredStyle.Render(countDown)
			row = lipgloss.JoinVertical(lipgloss.Left, row, countDown)
		} else {
			row = Style.PaddingRight(p.Pane.width - lipgloss.Width(countDown) - lipgloss.Width(row) - 2).Render(row)
			countDown = Theme.BlurredStyle.Render(countDown)
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, countDown)
		}

		if p.cursor == currentOrder {
			row = activeRowStyle.Width(p.Pane.width).Render(row)
		} else {
			row = inactiveRowStyle.Width(p.Pane.width).Render(row)
		}

		pendingOrderRows = append(pendingOrderRows, row)
		currentOrder++
	}

	pendingOrderContent := lipgloss.JoinVertical(lipgloss.Left, pendingOrderRows...)

	if len(pendingOrders) == 0 {
		pendingOrderContent = Theme.BlurredStyle.Render("No orders pending")
	}

	var execOrderRows []string

	execOrdersTitle := "Active"
	execOrdersTitleStyles := Style.Bold(true)
	execOrderRows = append(execOrderRows, execOrdersTitleStyles.Render(execOrdersTitle))

	for _, order := range executingOrders {

		var rows []string
		orderLabel := fmt.Sprintf("[%v] %v", order.GetStatus(), order.GetName())
		orderStyle := Style.Width(p.Pane.width).Align(lipgloss.Left)
		rows = append(rows, orderStyle.Render(orderLabel))

		barID := p.progressBars[order.GetID()]
		progressBar := PaneManager.Panes[barID]
		rows = append(rows, progressBar.View())
		orderContent := lipgloss.JoinVertical(lipgloss.Left, rows...)

		if p.cursor == currentOrder {
			orderContent = activeRowStyle.Width(p.Pane.width).Render(orderContent)
		} else {
			orderContent = inactiveRowStyle.Width(p.Pane.width).Render(orderContent)
		}

		execOrderRows = append(execOrderRows, orderContent)
		currentOrder++
	}

	execOrderContent := lipgloss.JoinVertical(lipgloss.Left, execOrderRows...)

	if len(executingOrders) == 0 {
		execOrderContent = Theme.BlurredStyle.Render("No orders queued")
	}

	var completedOrderRows []string

	completedOrdersTitle := "Completed"
	completedOrdersTitleStyles := Style.Bold(true)
	completedOrderRows = append(completedOrderRows, completedOrdersTitleStyles.Render(completedOrdersTitle))

	for _, order := range completedOrders {
		row := fmt.Sprintf("[%v] %v", order.GetStatus(), order.GetName())

		if p.cursor == currentOrder {
			row = activeRowStyle.Width(p.Pane.width).Render(row)
		} else {
			row = inactiveRowStyle.Width(p.Pane.width).Render(row)
		}

		completedOrderRows = append(completedOrderRows, row)

		currentOrder++
	}
	completedOrderContent := lipgloss.JoinVertical(lipgloss.Left, completedOrderRows...)

	if len(completedOrders) == 0 {
		completedOrderContent = Theme.BlurredStyle.Render("No orders completed")
	}

	content := lipgloss.JoinVertical(lipgloss.Left, title, pendingOrderContent, execOrderContent, completedOrderContent)

	return content
}

func (p *OrderStatusPane) initProgrssBars() {
	if p.progressBars == nil {
		p.progressBars = make(map[events.EventID]core.PaneID)
	}

	for _, order := range p.orderScheduler.PriorityQueue {
		if _, ok := p.progressBars[order.ID]; !ok {
			progressBar := NewLoadingBarPane("Order Status: "+order.GetName(), order.GetExecuteTick(), order.GetEndTick())
			id := PaneManager.AddPane(progressBar)
			p.progressBars[order.GetID()] = id

		}
	}
}
