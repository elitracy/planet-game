package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	. "github.com/elitracy/planets/models"
	. "github.com/elitracy/planets/state"
)

var (
	activeRowStyle = Style.
			Width(PaneManager.Width).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("212")).
			Padding(0, 1)

	inactiveRowStyle = Style.
				Width(PaneManager.Width).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(0, 1)
)

var (
	progressBarWidth int
)

type OrderStatusPane struct {
	Pane
	id             int
	title          string
	width          int
	height         int
	cursor         int
	orderScheduler *EventScheduler[Order]
	progressBars   map[Action]int
}

func NewOrderStatusPane(orderScheduler *EventScheduler[Order], title string) *OrderStatusPane {
	pane := &OrderStatusPane{
		title:          title,
		orderScheduler: orderScheduler,
	}

	return pane
}

func (p OrderStatusPane) GetId() int       { return p.id }
func (p *OrderStatusPane) SetId(id int)    { p.id = id }
func (p OrderStatusPane) GetTitle() string { return p.title }

func (p *OrderStatusPane) Init() tea.Cmd {
	p.updateProgressBars()
	return nil
}

func (p *OrderStatusPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	p.updateProgressBars()

	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.GetId() {
			p.width = msg.width - 2
			p.height = msg.height

			var cmds []tea.Cmd
			for _, val := range p.progressBars {
				cmds = append(cmds, paneResizeCmd(val, progressBarWidth, msg.height))
			}
			return p, tea.Batch(cmds...)

		}
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.orderScheduler.PriorityQueue)+len(State.CompletedOrders)-1 {
				p.cursor++
			}
		case "esc":
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

// TODO: refactor into switch function
func (p *OrderStatusPane) View() string {

	var pendingOrders []Order
	var executingOrders []Order
	var completedOrders []Order

	for _, order := range p.orderScheduler.PriorityQueue {
		if order.GetStatus() == Pending {
			pendingOrders = append(pendingOrders, order)
		}

		if order.GetStatus() == Executing {
			executingOrders = append(executingOrders, order)
		}
	}

	for _, order := range State.CompletedOrders {
		completedOrders = append(completedOrders, order)
	}

	p.updateProgressBars()

	title := Style.Width(p.width).AlignHorizontal(lipgloss.Center).Render(p.title + "\n")

	var pendingOrderRows []string

	pendingOrdersTitle := "Pending"
	pendingOrdersTitleStyles := Style.Bold(true)
	pendingOrderRows = append(pendingOrderRows, pendingOrdersTitleStyles.Render(pendingOrdersTitle))

	currentOrder := 0
	for _, order := range pendingOrders {
		row := fmt.Sprintf("[%v] %v", order.GetStatus(), order.GetName())

		countDown := fmt.Sprintf("ETA: %vs", (order.GetExecuteTick()-State.Tick)/TICKS_PER_SECOND)

		if lipgloss.Width(countDown)+lipgloss.Width(row) > p.width-5 {
			row = Style.Render(row)
			countDown = Theme.blurredStyle.Render(countDown)
			row = lipgloss.JoinVertical(lipgloss.Left, row, countDown)
		} else {
			row = Style.PaddingRight(p.width - lipgloss.Width(countDown) - lipgloss.Width(row) - 2).Render(row)
			countDown = Theme.blurredStyle.Render(countDown)
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, countDown)
		}

		if p.cursor == currentOrder && PaneManager.ActivePane().(Pane).GetId() == p.GetId() {
			row = activeRowStyle.Width(p.width).Render(row)
		} else {
			row = inactiveRowStyle.Width(p.width).Render(row)
		}

		pendingOrderRows = append(pendingOrderRows, row)
		currentOrder++
	}

	pendingOrderContent := lipgloss.JoinVertical(lipgloss.Left, pendingOrderRows...)

	if len(pendingOrders) == 0 {
		pendingOrderContent = Theme.blurredStyle.Render("No orders pending")
	}

	var execOrderRows []string

	execOrdersTitle := "Active"
	execOrdersTitleStyles := Style.Bold(true)
	execOrderRows = append(execOrderRows, execOrdersTitleStyles.Render(execOrdersTitle))

	for _, order := range executingOrders {

		var rows []string
		orderLabel := fmt.Sprintf("[%v] %v", order.GetStatus(), order.GetName())
		orderStyle := Style.Width(p.width).Align(lipgloss.Left)
		rows = append(rows, orderStyle.Render(orderLabel))

		for _, action := range order.GetActions() {
			progressBar := PaneManager.Panes[p.progressBars[action]]
			label := fmt.Sprintf("\n• [%v] %v", action.GetStatus(), action.GetDescription())

			label = Style.Width(lipgloss.Width(label)).Align(lipgloss.Left).Render(label)
			label = Style.PaddingRight(p.width - lipgloss.Width(label) - lipgloss.Width(progressBar.View()) - 5).Render(label)

			var row string
			if lipgloss.Width(label)+lipgloss.Width(progressBar.View()) > p.width-5 {
				row = lipgloss.JoinVertical(lipgloss.Left, label, progressBar.View())
			} else {
				row = lipgloss.JoinHorizontal(lipgloss.Bottom, label, progressBar.View())
			}

			rows = append(rows, row)
		}

		orderContent := lipgloss.JoinVertical(lipgloss.Left, rows...)

		if p.cursor == currentOrder && PaneManager.ActivePane().(Pane).GetId() == p.GetId() {
			orderContent = activeRowStyle.Width(p.width).Render(orderContent)
		} else {
			orderContent = inactiveRowStyle.Width(p.width).Render(orderContent)
		}

		execOrderRows = append(execOrderRows, orderContent)
		currentOrder++
	}

	execOrderContent := lipgloss.JoinVertical(lipgloss.Left, execOrderRows...)

	if len(executingOrders) == 0 {
		execOrderContent = Theme.blurredStyle.Render("No orders queued")
	}

	var completedOrderRows []string

	completedOrdersTitle := "Completed"
	completedOrdersTitleStyles := Style.Bold(true)
	completedOrderRows = append(completedOrderRows, completedOrdersTitleStyles.Render(completedOrdersTitle))

	for _, order := range completedOrders {
		row := fmt.Sprintf("[%v] %v", order.GetStatus(), order.GetName())

		if p.cursor == currentOrder && PaneManager.ActivePane().(Pane).GetId() == p.GetId() {
			row = activeRowStyle.Width(p.width).Render(row)
		} else {
			row = inactiveRowStyle.Width(p.width).Render(row)
		}

		completedOrderRows = append(completedOrderRows, row)

		currentOrder++
	}
	completedOrderContent := lipgloss.JoinVertical(lipgloss.Left, completedOrderRows...)

	if len(completedOrders) == 0 {
		completedOrderContent = Theme.blurredStyle.Render("No orders completed")
	}

	content := lipgloss.JoinVertical(lipgloss.Left, title, pendingOrderContent, execOrderContent, completedOrderContent)

	return content
}

func (p *OrderStatusPane) updateProgressBars() {
	if p.progressBars == nil {
		p.progressBars = make(map[Action]int)
	}

	for _, order := range p.orderScheduler.PriorityQueue {
		for _, action := range order.GetActions() {
			if _, ok := p.progressBars[action]; !ok {
				progressBar := NewLoadingBarPane("Order Status: "+order.GetName(), action.GetExecuteTick(), action.GetExecuteTick()+action.GetDuration())
				id := PaneManager.AddPane(progressBar)
				p.progressBars[action] = id
			}

		}
	}
}
