package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	. "github.com/elitracy/planets/models"
)

var (
	activeRowStyle = lipgloss.NewStyle().
			Width(PaneManager.Width).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("212")).
			Padding(0, 1)

	inactiveRowStyle = lipgloss.NewStyle().
				Width(PaneManager.Width).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(0, 1)
)

type OrderStatusPane struct {
	Pane
	id             int
	title          string
	orderScheduler *EventScheduler[*Order]
	cursor         int
}

func NewOrderStatusPane(orderScheduler *EventScheduler[*Order], title string) *OrderStatusPane {
	pane := &OrderStatusPane{
		title:          title,
		orderScheduler: *&orderScheduler,
	}

	return pane
}

func (p OrderStatusPane) GetId() int       { return p.id }
func (p *OrderStatusPane) SetId(id int)    { p.id = id }
func (p OrderStatusPane) GetTitle() string { return p.title }

func (p *OrderStatusPane) Init() tea.Cmd { return nil }

func (p *OrderStatusPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.orderScheduler.PriorityQueue)-1 {
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

func (p *OrderStatusPane) View() string {
	title := p.title + "\n"
	content := ""

	for i, order := range p.orderScheduler.PriorityQueue {
		orderContent := ""
		orderContent += fmt.Sprintf("[%v] %v %v", order.Status, order.Type, order.TargetEntity.GetName())

		for _, action := range order.Actions {
			orderContent += fmt.Sprintf("\n • [%v] %v", action.Status, action.Type)
		}

		if p.cursor == i {
			orderContent = activeRowStyle.Render(orderContent)
		} else {
			orderContent = inactiveRowStyle.Render(orderContent)
		}

		content += orderContent + "\n"
	}

	content = lipgloss.JoinVertical(lipgloss.Center, title, content)

	return content
}
