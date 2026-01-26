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

type OrderListPane struct {
	*Pane
	cursor       int
	orders       []*orders.Order
	progressBars map[events.EventID]core.PaneID
	status       events.EventStatus
	theme        UITheme
}

func NewOrderListPane(orders []*orders.Order, status events.EventStatus) *OrderListPane {
	pane := &OrderListPane{
		Pane:   &Pane{},
		orders: orders,
		status: status,
	}

	return pane
}

func (p *OrderListPane) Init() tea.Cmd {
	p.initProgrssBars()
	return nil
}

func (p *OrderListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.orders) {
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

func (p *OrderListPane) View() string {
	p.theme = GetPaneTheme(p)

	var rows []string

	title := fmt.Sprintf("%v Orders", p.status)
	titleStyled := Style.Width(p.width).Bold(true).Align(lipgloss.Center).PaddingBottom(1).Render(title)

	for i, order := range p.orders {

		if order.Status != p.status {
			continue
		}

		row := order.GetName()
		if order.Status == events.EventExecuting {
			countDown := fmt.Sprintf("ETA: %vs", int((order.GetEndTick()-state.State.CurrentTick)/consts.TICKS_PER_SECOND))
			progressBar := PaneManager.Panes[p.progressBars[order.GetID()]]
			row += fmt.Sprintf(" %v %v", progressBar.View(), countDown)
		}

		if p.cursor == i {
			row = p.theme.FocusedStyle.Render(row)
		} else {
			row = p.theme.BlurredStyle.Render(row)
		}

		rows = append(rows, row)
	}

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, rows...)

	content := lipgloss.JoinVertical(lipgloss.Left, titleStyled, infoContainer)

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
