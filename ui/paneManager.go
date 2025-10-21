package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"

	. "github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/interfaces"
	"github.com/elitracy/planets/core/logging"
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
	id     int
	title  string
	width  int
	height int

	FocusStack FocusStack
	Panes      map[int]tea.Model
	currentID  int
	UITick     Tick
}

var PaneManager = NewPaneManager()

func (p paneManager) GetId() int       { return p.id }
func (p *paneManager) SetId(id int)    { p.id = id }
func (p paneManager) GetTitle() string { return p.title }
func (p paneManager) GetWidth() int    { return p.width }
func (p paneManager) GetHeight() int   { return p.height }
func (p *paneManager) SetWidth(w int)  { p.width = w }
func (p *paneManager) SetHeight(h int) {
	logging.Info("TERM HEIGHT: %v", h)
	p.height = h
}

func NewPaneManager() *paneManager {

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		logging.Error("Failed to intialize Pane Manager: %v", err)
		return nil
	}

	return &paneManager{
		FocusStack: FocusStack{},
		Panes:      make(map[int]tea.Model, 0),
		currentID:  0,
		UITick:     0,
		width:      width,
		height:     height,
	}
}

func (p *paneManager) Init() tea.Cmd {
	var cmds []tea.Cmd

	for i := range p.Panes {
		cmds = append(cmds, p.Panes[i].Init())
		cmds = append(cmds, paneResizeCmd(i, p.width, p.height))
	}
	cmds = append(cmds, TickCmd(p.UITick))

	return tea.Batch(cmds...)
}

func (p *paneManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case pushFocusMsg:
		p.PushFocusStack(msg.id)
	case popFocusMsg:
		p.PopFocusStack()
	case TickMsg:
		p.UITick++
		cmds = append(cmds, TickCmd(p.UITick))
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
	}

	for id := range p.Panes {
		var cmd tea.Cmd
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			cmds = append(cmds, paneResizeCmd(id, p.width, p.height))
		case tea.KeyMsg:
			if p.Panes[id].(Pane).GetId() == p.ActivePane().(Pane).GetId() {
				p.Panes[id], cmd = p.Panes[id].Update(msg)
			}
		default:
			p.Panes[id], cmd = p.Panes[id].Update(msg)
		}

		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return p, tea.Batch(cmds...)
}

func (p *paneManager) View() string { return p.ActivePane().View() }

func (p *paneManager) AddPane(pane Pane) int {
	p.currentID++
	pane.SetId(p.currentID)
	p.Panes[p.currentID] = pane.(tea.Model)

	if model, ok := pane.(tea.Model); ok {
		model.Init()
	} else {
		logging.Error("Pane[%v] is not a model", pane.GetId())
	}

	return p.currentID
}

func (p *paneManager) RemovePane(id int) {
	delete(p.Panes, id)
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
