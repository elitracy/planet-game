package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
)

type pushFocusMsg struct{ id core.PaneID }
type popFocusMsg struct{ id core.PaneID }

type focusTabsMsg struct{ lastActiveID core.PaneID }

type paneResizeMsg struct {
	paneID core.PaneID
	width  int
	height int
}

func pushFocusCmd(id core.PaneID) tea.Cmd { return func() tea.Msg { return pushFocusMsg{id} } }
func popFocusCmd(id core.PaneID) tea.Cmd  { return func() tea.Msg { return popFocusMsg{id} } }

func focusTabsCmd(lastActiveID core.PaneID) tea.Cmd {
	return func() tea.Msg { return focusTabsMsg{lastActiveID: lastActiveID} }
}

func paneResizeCmd(id core.PaneID, width, height int) tea.Cmd {
	return func() tea.Msg { return paneResizeMsg{paneID: id, width: width, height: height} }
}

type paneManager struct {
	*Pane

	FocusStack  FocusStack
	Panes       map[core.PaneID]ManagedPane
	Root        *RootPane
	focusedTabs bool
	currentID   core.PaneID
	UITick      core.Tick
}

var PaneManager = NewPaneManager()

func NewPaneManager() *paneManager {

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		logging.Error("Failed to intialize Pane Manager: %v", err)
		return nil
	}

	pm := &paneManager{
		FocusStack: FocusStack{},
		Panes:      make(map[core.PaneID]ManagedPane),
		currentID:  0,
		UITick:     0,
		Pane: &Pane{
			width:  width,
			height: height,
		},
	}

	return pm
}

func (p *paneManager) Init() tea.Cmd {
	var cmds []tea.Cmd

	for i := range p.Panes {
		cmds = append(cmds, p.Panes[i].Init())
		cmds = append(cmds, paneResizeCmd(i, p.Pane.width, p.Pane.height))
	}
	cmds = append(cmds, core.TickCmd(p.UITick))

	cmds = append(cmds, focusTabsCmd(-1))

	return tea.Batch(cmds...)
}

func (p *paneManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case focusTabsMsg:
		p.focusedTabs = true
	case pushFocusMsg:
		p.focusedTabs = false
		p.PushFocusStack(msg.id)
	case popFocusMsg:
		lastPane := p.ActivePane()

		p.PopFocusStack()

		if len(p.FocusStack.stack) < 2 {
			cmds = append(cmds, focusTabsCmd(lastPane.ID()))
		}
	case tea.KeyMsg:
		var cmd tea.Cmd
		var model tea.Model

		if p.focusedTabs {
			model, cmd = p.Panes[p.Root.Pane.id].Update(msg)
			p.Panes[p.Root.Pane.id] = model.(ManagedPane)
		}
		cmds = append(cmds, cmd)
	case core.TickMsg:
		p.UITick++
		cmds = append(cmds, core.TickCmd(p.UITick))
	case tea.WindowSizeMsg:
		p.Pane.width = msg.Width
		p.Pane.height = msg.Height
	}

	for id := range p.Panes {
		var cmd tea.Cmd
		var model tea.Model

		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			cmds = append(cmds, paneResizeCmd(id, p.Pane.width, p.Pane.height))
		case tea.KeyMsg:
			if !p.focusedTabs && p.Panes[id].ID() == p.ActivePane().ID() {
				model, cmd = p.Panes[id].Update(msg)
				p.Panes[id] = model.(ManagedPane)
			}
		default:
			model, cmd = p.Panes[id].Update(msg)

			p.Panes[id] = model.(ManagedPane)
		}

		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return p, tea.Batch(cmds...)
}

func (p *paneManager) View() string { return p.Root.View() }

func (p *paneManager) AddPane(pane ManagedPane) core.PaneID {
	p.currentID++

	pane.SetID(p.currentID)
	p.Panes[p.currentID] = pane

	logging.Info("%v: %v", p.currentID, pane.Title())
	p.Panes[p.currentID].Init()

	return p.currentID
}

func (p *paneManager) RemovePane(id core.PaneID) { delete(p.Panes, id) }

type FocusStack struct{ stack []ManagedPane }

func (p *paneManager) PushFocusStack(id core.PaneID) {
	p.FocusStack.stack = append(p.FocusStack.stack, p.Panes[id])
}

func (p *paneManager) PopFocusStack() ManagedPane {
	if len(p.FocusStack.stack) > 1 {
		p.FocusStack.stack = p.FocusStack.stack[:len(p.FocusStack.stack)-1]
	}

	pane := p.FocusStack.stack[len(p.FocusStack.stack)-1]
	return pane
}

func (p paneManager) ActivePane() ManagedPane {
	if len(p.FocusStack.stack) < 2 {
		return p.FocusStack.stack[0]
	}

	return p.FocusStack.stack[len(p.FocusStack.stack)-1]
}
