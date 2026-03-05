package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/orders"
)

var orderStatusTypes = []task.Status{
	task.Pending,
	task.Executing,
	task.Complete,
	task.Failed,
}

type OrderStatusListPane struct {
	*engine.Pane
	cursor         int
	orderScheduler *task.TaskScheduler[*orders.Order]
	orderList *engine.LayoutNode
	theme          UITheme
}

func NewOrderStatusListPane(title string, orderScheduler *task.TaskScheduler[*orders.Order]) *OrderStatusListPane {
	pane := &OrderStatusListPane{
		Pane:           engine.NewPane(title, engine.NewKeyBindings()),
		cursor:         0,
		orderScheduler: orderScheduler,
	}

	return pane
}

func (p *OrderStatusListPane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Select, "enter").
		Set(engine.Back, "esc").
		Set(engine.Up, "k").
		Set(engine.Down, "j")

	orderList := NewOrderListPane(task.Status(p.cursor))
	PaneManager.AddPane(orderList)
	layout := engine.NewLayoutNode(orderList, engine.LayoutHorizontal, 0.75)
	p.orderList = layout


	return tea.Sequence(popLayoutCmd(), pushLayoutCmd(layout, p.ID()))
}

func (p *OrderStatusListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case p.GetKeys().Get(engine.Up):
			if p.cursor > 0 {
				p.cursor--
			}

			p.orderList.Pane.(*OrderListPane).Status = task.Status(p.cursor)
		case p.GetKeys().Get(engine.Down):
			if p.cursor < len(orderStatusTypes)-1 {
				p.cursor++
			}

			engine.Info("order list pane: %v", p.orderList)
			p.orderList.Pane.(*OrderListPane).Status = task.Status(p.cursor)
		case p.GetKeys().Get(engine.Select):
			return p, tea.Sequence(pushFocusStackCmd(p.orderList.Pane.ID()))
		case p.GetKeys().Get(engine.Back):
			return p, tea.Sequence(popFocusStackCmd())
		case p.GetKeys().Get(engine.Quit):
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

		row := task.Status(status).String()

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
	titleStyled := Style.Width(p.Width()).Bold(true).AlignHorizontal(lipgloss.Center).PaddingBottom(1).Render(title)

	content := lipgloss.JoinVertical(lipgloss.Left, titleStyled, infoContainer)

	content = Style.Height(p.Height()).Width(p.Width()).Render(content)

	return content
}
