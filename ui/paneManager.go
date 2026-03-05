package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/config"
	"golang.org/x/term"
)

type paneManager struct {
	*engine.Pane

	TabLine    engine.ManagedPane
	StatusLine engine.ManagedPane
	layout     *engine.LayoutNode

	Panes      map[engine.PaneID]engine.ManagedPane
	currentID  engine.PaneID
	focusStack []engine.PaneID
	state      *game.GameState

	CurrentUITick engine.Tick
}

var PaneManager *paneManager

func InitPaneManager() { PaneManager = NewPaneManager() }

func NewPaneManager() *paneManager {

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Pane Manager: %v", err))
	}

	pm := &paneManager{
		Panes:         make(map[engine.PaneID]engine.ManagedPane),
		currentID:     0,
		CurrentUITick: 0,
		TabLine:       NewTablinePane(),
		StatusLine:    NewStatusLinePane(game.State.CurrentTick),
		Pane:          engine.NewPane("Pane Manager", nil),
		state:         game.State,
	}

	pm.SetSize(width, height)

	pm.Panes[-1] = NewErrorPane("No content.")

	return pm
}

func (p *paneManager) PushFocusStack(id engine.PaneID) {
	p.focusStack = append(p.focusStack, id)
	pane := p.Panes[id]
	p.StatusLine.(*StatusLinePane).SetKeys(pane.GetKeys())
}

func (p *paneManager) PopFocusStack() {
	if len(p.focusStack) <= 1 {
		return
	}

	p.focusStack = p.focusStack[:len(p.focusStack)-1]

	paneID := p.focusStack[len(p.focusStack)-1]
	p.StatusLine.(*StatusLinePane).SetKeys(p.Panes[paneID].GetKeys())
}

func (p *paneManager) PeekFocusStack() engine.PaneID {
	if len(p.focusStack) <= 0 {
		return -1
	}

	return p.focusStack[len(p.focusStack)-1]
}

func (p *paneManager) AddTab(layout *engine.LayoutNode) {
	p.TabLine.(*TabLinePane).tabs = append(p.TabLine.(*TabLinePane).tabs, layout)

}

func (p *paneManager) Init() tea.Cmd {

	var cmds []tea.Cmd

	if p.TabLine != nil {
		cmds = append(cmds, p.TabLine.Init())
	}

	if p.StatusLine != nil {
		cmds = append(cmds, p.StatusLine.Init())
	}

	firstTab := p.TabLine.(*TabLinePane).tabs[0]
	p.layout = firstTab
	p.PushFocusStack(firstTab.Pane.ID())

	cmds = append(
		cmds,
		engine.TickCmd(p.state.CurrentTick),
		config.UITickCmd(p.CurrentUITick),
	)

	return tea.Batch(cmds...)
}

func (p *paneManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case setLayoutMsg:
		if p.layout.Pane != nil {
			cmd := p.layout.Pane.Init()
			cmds = append(cmds, cmd)
		}
		p.layout = msg.layout
	case pushLayoutMsg:
		if p.layout.Pane != nil {
			cmd := p.layout.Pane.Init()
			cmds = append(cmds, cmd)
		}
		p.layout.Push(msg.layout, msg.targetPaneID)
	case popLayoutMsg:
		p.layout.Pop(p.PeekFocusStack())
	case pushFocusStackMsg:
		p.PushFocusStack(msg.id)
	case popFocusStackMsg:
		p.PopFocusStack()
	case flushFocusStackMsg:
		p.focusStack = nil
	case tea.WindowSizeMsg:
		p.SetSize(msg.Width, msg.Height)
	case engine.TickMsg:
		cmds = []tea.Cmd{engine.TickCmd(p.state.CurrentTick)}
		for id, pane := range p.Panes {
			model, cmd := pane.Update(msg)
			p.Panes[id] = model.(engine.ManagedPane)

			cmds = append(cmds, cmd)
		}

		if p.TabLine != nil {
			model, cmd := p.TabLine.Update(msg)
			p.TabLine = model.(*TabLinePane)
			cmds = append(cmds, cmd)
		}

		if p.StatusLine != nil {
			model, cmd := p.StatusLine.Update(msg)
			p.StatusLine = model.(*StatusLinePane)
			cmds = append(cmds, cmd)
		}

	case config.UITickMsg:
		p.CurrentUITick++
		cmds = []tea.Cmd{config.UITickCmd(p.CurrentUITick)}
		for id, pane := range p.Panes {
			model, cmd := pane.Update(msg)
			p.Panes[id] = model.(engine.ManagedPane)

			cmds = append(cmds, cmd)
		}

		if p.TabLine != nil {
			model, cmd := p.TabLine.Update(msg)
			p.TabLine = model.(*TabLinePane)
			cmds = append(cmds, cmd)
		}

		if p.StatusLine != nil {
			model, cmd := p.StatusLine.Update(msg)
			p.StatusLine = model.(*StatusLinePane)
			cmds = append(cmds, cmd)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab":
			_, cmd := p.TabLine.Update(msg)
			return p, cmd
		}

		if pane, ok := p.Panes[p.PeekFocusStack()]; ok {
			model, cmd := pane.Update(msg)
			p.Panes[p.PeekFocusStack()] = model.(engine.ManagedPane)
			return p, cmd
		}
	default:
		for id, pane := range p.Panes {
			model, cmd := pane.Update(msg)
			p.Panes[id] = model.(engine.ManagedPane)
			cmds = append(cmds, cmd)
		}
	}

	return p, tea.Batch(cmds...)
}

func (p *paneManager) View() string {
	mainContent := "No tab selected"

	if len(p.layout.Children) == 0 && p.layout.Pane == nil {
		return mainContent
	}

	tabLineStyle := Style.Width(p.Width()).Border(lipgloss.NormalBorder(), false, false, true, false).Render(p.TabLine.View())
	statusLineStyle := Style.Width(p.Width()).Border(lipgloss.NormalBorder(), true, false, false, false).Render(p.StatusLine.View())

	contentHeight := p.Height() - lipgloss.Height(tabLineStyle) - lipgloss.Height(statusLineStyle)
	contentView := p.layout.Render(p.Width(), contentHeight)

	return lipgloss.JoinVertical(lipgloss.Top, tabLineStyle, contentView, statusLineStyle)
}

func (p *paneManager) AddPane(pane engine.ManagedPane) engine.PaneID {
	p.currentID++
	id := p.currentID
	pane.SetID(id)
	p.Panes[id] = pane

	return id
}

func (p *paneManager) RemovePane(id engine.PaneID) { delete(p.Panes, id) }
