package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/orders"
)

var orderStatusTypes = []events.EventStatus{
	events.EventPending,
	events.EventExecuting,
	events.EventComplete,
	events.EventFailed,
}

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
		Set(Select, "enter").
		Set(Back, "esc").
		Set(Up, "k").
		Set(Down, "j")

	orderList := NewOrderListPane(events.EventStatus(p.cursor))
	paneID := PaneManager.AddPane(orderList)
	return tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
}

func (p *OrderStatusListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

			orderList := NewOrderListPane(events.EventStatus(p.cursor))
			paneID := PaneManager.AddPane(orderList)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case p.keys.Get(Down):
			if p.cursor < len(orderStatusTypes)-1 {
				p.cursor++
			}

			orderList := NewOrderListPane(events.EventStatus(p.cursor))
			paneID := PaneManager.AddPane(orderList)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case p.keys.Get(Select):
			return p, tea.Sequence(pushFocusStackCmd(PaneManager.PeekDetailPaneStack().ID()))
		case p.keys.Get(Back):
			return p, tea.Sequence(popFocusStackCmd())
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
	for status := range orderStatusTypes {

		row := events.EventStatus(status).String()

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
