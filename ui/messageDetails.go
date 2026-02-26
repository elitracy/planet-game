package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/events"
)

type MessageDetailsPane struct {
	*engine.Pane
	message *events.Event

	theme     UITheme
	cursor    int
	infoTable engine.ManagedPane
}

func (p *MessageDetailsPane) Init() tea.Cmd {
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Back, "esc").
		Set(engine.Up, "k").
		Set(engine.Down, "j")

	keymaps := make(map[string]func() tea.Cmd)
	keymaps[p.GetKeys().Get(engine.Back)] = func() tea.Cmd {
		return tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
	}

	infoTable := p.createInfoTable()
	p.infoTable = NewInfoTablePane(
		infoTable,
		keymaps,
	)

	PaneManager.AddPane(p.infoTable)

	return nil
}

func (p *MessageDetailsPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case paneResizeMsg:
		p.SetSize(msg.width, msg.height)
	case engine.TickMsg:
		p.infoTable.(*InfoTablePane).SetTheme(GetPaneTheme(p))
	case config.UITickMsg:
		p.infoTable.(*InfoTablePane).table.SetRows(p.createRows())
	case tea.KeyMsg:
		switch msg.String() {
		case p.GetKeys().Get(engine.Up):
			if p.cursor > 0 {
				p.cursor--
			}
			engine.Info("up")
		case p.GetKeys().Get(engine.Down):
			if p.cursor < len(p.infoTable.(*InfoTablePane).table.Rows()) {
				p.cursor++
			}
			engine.Info("down")
		case p.GetKeys().Get(engine.Back):
			return p, tea.Sequence(popFocusStackCmd(), popDetailStackCmd())
		case p.GetKeys().Get(engine.Quit):
			return p, tea.Quit
		}
	}

	model, cmd := p.infoTable.Update(msg)
	cmds = append(cmds, cmd)
	p.infoTable = model.(engine.ManagedPane)

	return p, tea.Batch(cmds...)

}

func (p *MessageDetailsPane) View() string {
	p.theme = GetPaneTheme(p)

	title := p.Title()
	titleStyled := Style.Width(p.Width()).AlignHorizontal(lipgloss.Center).Bold(true).PaddingBottom(1).Render(title)

	return lipgloss.JoinVertical(lipgloss.Left, titleStyled, p.infoTable.View())
}

func (p MessageDetailsPane) createInfoTable() table.Model {
	infoTable := table.New(
		table.WithColumns(p.createColumns()),
		table.WithRows(p.createRows()),
		table.WithFocused(true),
		table.WithHeight(40),
	)

	return infoTable
}

func (p *MessageDetailsPane) createColumns() []table.Column {

	columns := []table.Column{
		{Title: "Title", Width: 15},
		{Title: "Body", Width: 35},
		{Title: "Severity", Width: 10},
		{Title: "Target", Width: 20},
		{Title: "Time Remaining", Width: 20},
	}
	return columns
}

func (p *MessageDetailsPane) createRows() []table.Row {

	rows := []table.Row{}

	// for _, event := range p..ActiveEvents {
	// 	row := table.Row{event.Description}
	// 	rows = append(rows, row)
	// }

	return rows
}
