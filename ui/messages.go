package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/events"
)

type MessagePane struct {
	*engine.Pane

	cursor         int
	events         *events.EventManager
	eventInfoTable engine.ManagedPane
	theme          UITheme
}

func NewMessagePane(title string, events *events.EventManager) *MessagePane {
	pane := &MessagePane{
		Pane:   engine.NewPane(title, engine.NewKeyBindings()),
		events: events,
	}

	return pane
}

func (p *MessagePane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Back, "esc").
		Set(engine.Up, "k").
		Set(engine.Down, "j")

	keymaps := make(map[string]func() tea.Cmd)
	keymaps[p.GetKeys().Get(engine.Back)] = func() tea.Cmd {
		return tea.Sequence(popLayoutCmd(), popFocusStackCmd())
	}

	infoTable := p.createInfoTable()
	p.eventInfoTable = NewInfoTablePane(
		infoTable,
		keymaps,
	)

	PaneManager.AddPane(p.eventInfoTable)

	return nil
}

func (p *MessagePane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case engine.TickMsg:
		p.eventInfoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case config.UITickMsg:
		p.eventInfoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case tea.KeyMsg:
		switch msg.String() {
		case p.GetKeys().Get(engine.Up):
			if p.cursor > 0 {
				p.cursor--
			}
			engine.Info("up")
		case p.GetKeys().Get(engine.Down):
			if p.cursor < len(p.events.ActiveEvents) {
				p.cursor++
			}
			engine.Info("down")
		case p.GetKeys().Get(engine.Back):
			return p, tea.Sequence(popFocusStackCmd(), popLayoutCmd())
		case p.GetKeys().Get(engine.Quit):
			return p, tea.Quit
		}
	}

	model, cmd := p.eventInfoTable.Update(msg)
	cmds = append(cmds, cmd)
	p.eventInfoTable = model.(engine.ManagedPane)

	return p, tea.Batch(cmds...)

}

func (p *MessagePane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.Title()
	titleStyled := Style.Width(p.Width()).AlignHorizontal(lipgloss.Center).Bold(true).PaddingBottom(1).Render(title)

	return lipgloss.JoinVertical(lipgloss.Left, titleStyled, p.eventInfoTable.View())
}

func (p MessagePane) createInfoTable() table.Model {
	infoTable := table.New(
		table.WithColumns(p.createColumns()),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(40),
	)

	return infoTable
}

func (p *MessagePane) createColumns() []table.Column {

	columns := []table.Column{
		{Title: "", Width: 35},
	}
	return columns
}

func (p *MessagePane) createRows() []table.Row {

	rows := []table.Row{}

	for _, event := range p.events.ActiveEvents {
		row := table.Row{event.Description}
		rows = append(rows, row)
	}

	return rows
}
