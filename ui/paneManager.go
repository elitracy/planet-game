package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	. "github.com/elitracy/planets/models"
)

type pushFocusMsg struct{ id int }
type popFocusMsg struct{}

type paneResizeMsg struct {
	paneID int
	width  int
	height int
}

func pushFocusCmd(id int) tea.Cmd { return func() tea.Msg { return pushFocusMsg{id} } }
func popFocusCmd() tea.Cmd        { return func() tea.Msg { return popFocusMsg{} } }

func paneResizeCmd(id, width, height int) tea.Cmd {
	return func() tea.Msg { return paneResizeMsg{paneID: id, width: width, height: height} }
}

type paneManager struct {
	FocusStack FocusStack
	Panes      map[int]tea.Model
	currentID  int
	UITick     int
	Width      int
	Height     int

	id    int
	title string
}

func NewPaneManager() paneManager {
	return paneManager{
		FocusStack: FocusStack{},
		Panes:      make(map[int]tea.Model, 0),
		currentID:  0,
		UITick:     0,
	}
}

var PaneManager = NewPaneManager()

func (p *paneManager) AddPane(pane Pane) int {
	p.currentID++
	pane.SetId(p.currentID)
	p.Panes[p.currentID] = pane.(tea.Model)

	return p.currentID
}

func (p *paneManager) RemovePane(id int) {
	delete(p.Panes, id)
}

func (p paneManager) GetId() int       { return p.id }
func (p *paneManager) SetId(id int)    { p.id = id }
func (p paneManager) GetTitle() string { return p.title }

func (p *paneManager) Init() tea.Cmd {
	var cmds []tea.Cmd

	for i := range p.Panes {
		cmds = append(cmds, p.Panes[i].Init())
	}
	cmds = append(cmds, tick(p.UITick))

	return tea.Batch(cmds...)
}

func (p *paneManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case pushFocusMsg:
		p.PushFocusStack(msg.id)
	case popFocusMsg:
		p.PopFocusStack()
	case tickMsg:
		p.UITick++
		cmds = append(cmds, tick(p.UITick))
	case tea.WindowSizeMsg:
		p.Width = msg.Width - 10
		p.Height = msg.Height - 10
	}

	for key := range p.Panes {
		var cmd tea.Cmd
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if p.Panes[key].(Pane).GetId() == p.ActivePane().(Pane).GetId() {
				p.Panes[key], cmd = p.Panes[key].Update(msg)
			}
		default:
			p.Panes[key], cmd = p.Panes[key].Update(msg)
		}

		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return p, tea.Batch(cmds...)

}

func (p *paneManager) View() string {
	viewStyle := Style.Padding(2)
	view := viewStyle.Render(p.ActivePane().View())

	return view
}

type FocusStack struct {
	stack []tea.Model
}

func (p *paneManager) PushFocusStack(id int) {
	p.FocusStack.stack = append(p.FocusStack.stack, p.Panes[id])
}

func (p *paneManager) PopFocusStack() tea.Model {

	if len(p.FocusStack.stack) > 1 {
		p.FocusStack.stack = p.FocusStack.stack[:len(p.FocusStack.stack)-1]
	}

	pane := p.FocusStack.stack[len(p.FocusStack.stack)-1]
	return pane
}

func (p paneManager) ActivePane() tea.Model {
	if len(p.FocusStack.stack) < 2 {
		return p.FocusStack.stack[0]
	}

	return p.FocusStack.stack[len(p.FocusStack.stack)-1]
}
