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

var mainWidthPercentage float32

var mainWidth int
var mainHeight int

var detailWidth int
var detailHeight int

type paneManager struct {
	*engine.Pane

	TabLine         engine.ManagedPane
	StatusLine      engine.ManagedPane
	MainPane        engine.ManagedPane
	DetailPaneStack []engine.ManagedPane

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
		TabLine:       NewTablinePane([]engine.ManagedPane{}),
		StatusLine:    NewStatusLinePane(game.State.CurrentTick),
		Pane:          engine.NewPane("Pane Manager", nil),
		state:         game.State,
	}

	pm.SetSize(width, height)

	mainWidthPercentage = .25

	mainWidth = int(float32(pm.Width()) * mainWidthPercentage)
	detailWidth = int(float32(pm.Width()) * (1 - mainWidthPercentage))

	mainHeight = pm.Height()
	detailHeight = pm.Height()

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

func (p *paneManager) AddTab(pane engine.ManagedPane) {
	p.TabLine.(*TabLinePane).tabs = append(p.TabLine.(*TabLinePane).tabs, pane)
}

func (p *paneManager) SetMainPane(pane engine.ManagedPane) {
	p.MainPane = pane
}

func (p *paneManager) PushDetailPaneStack(pane engine.ManagedPane) {
	p.DetailPaneStack = append(p.DetailPaneStack, pane)
}

func (p *paneManager) PopDetailPaneStack() {
	if len(p.DetailPaneStack) <= 0 {
		return
	}

	pane := p.DetailPaneStack[len(p.DetailPaneStack)-1]
	p.DetailPaneStack = p.DetailPaneStack[:len(p.DetailPaneStack)-1]

	p.RemovePane(pane.ID())
}

func (p *paneManager) PeekDetailPaneStack() engine.ManagedPane {
	if len(p.DetailPaneStack) <= 0 {
		return NewErrorPane("No detail selected")
	}

	return p.DetailPaneStack[len(p.DetailPaneStack)-1]
}

func (p *paneManager) FlushDetailPaneStack() {
	p.DetailPaneStack = nil
}

func (p *paneManager) Init() tea.Cmd {

	var cmds []tea.Cmd

	if p.MainPane == nil {
		p.SetMainPane(p.TabLine.(*TabLinePane).tabs[0])
	}

	if p.TabLine != nil {
		cmds = append(cmds, p.TabLine.Init())
	}

	if p.StatusLine != nil {
		cmds = append(cmds, p.StatusLine.Init())
	}

	if p.MainPane != nil {
		cmds = append(cmds, p.MainPane.Init())
	}

	if p.PeekDetailPaneStack() != nil {
		cmds = append(cmds, p.PeekDetailPaneStack().Init())
	}

	cmds = append(
		cmds,
		paneResizeCmd(p.MainPane.ID(), mainWidth, mainHeight),
		paneResizeCmd(p.PeekDetailPaneStack().ID(), detailWidth, detailHeight),
		engine.TickCmd(p.state.CurrentTick),
		config.UITickCmd(p.CurrentUITick),
	)

	p.PushFocusStack(p.MainPane.ID())

	return tea.Sequence(cmds...)
}

func (p *paneManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if p.MainPane == nil {
		p.MainPane = NewErrorPane("No content selected")
	}

	switch msg := msg.(type) {
	case paneResizeMsg:
		if pane, ok := p.Panes[msg.paneID]; ok {
			model, _ := pane.Update(msg)
			p.Panes[msg.paneID] = model.(engine.ManagedPane)
		}
	case setMainFocusMsg:
		if pane, ok := p.Panes[msg.id]; ok {
			p.SetMainPane(pane)
			return p, tea.Sequence(paneResizeCmd(pane.ID(), mainWidth, mainHeight), p.MainPane.Init())
		}
	case pushDetailStackMsg:
		if pane, ok := p.Panes[msg.id]; ok {
			pane.SetSize(detailWidth, detailHeight)
			p.PushDetailPaneStack(pane)
			return p, tea.Sequence(paneResizeCmd(p.PeekDetailPaneStack().ID(), detailWidth, detailHeight))
		}
	case popDetailStackMsg:
		p.PopDetailPaneStack()
		return p, tea.Sequence(paneResizeCmd(p.PeekDetailPaneStack().ID(), detailWidth, detailHeight))
	case flushDetailStackMsg:
		p.FlushDetailPaneStack()
	case pushFocusStackMsg:
		p.PushFocusStack(msg.id)
	case popFocusStackMsg:
		p.PopFocusStack()
	case flushFocusStackMsg:
		p.focusStack = nil
	case tea.WindowSizeMsg:
		p.SetSize(msg.Width, msg.Height)
	case engine.TickMsg:
		cmds := []tea.Cmd{engine.TickCmd(p.state.CurrentTick)}
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

		return p, tea.Batch(cmds...)

	case config.UITickMsg:
		p.CurrentUITick++
		cmds := []tea.Cmd{config.UITickCmd(p.CurrentUITick)}
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

		return p, tea.Batch(cmds...)
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
		var cmds []tea.Cmd
		for id, pane := range p.Panes {
			model, cmd := pane.Update(msg)
			p.Panes[id] = model.(engine.ManagedPane)
			cmds = append(cmds, cmd)
		}
		return p, tea.Batch(cmds...)
	}

	return p, nil
}

func (p *paneManager) View() string {
	mainContent := "No tab selected"
	if p.MainPane != nil {
		mainContent = p.MainPane.View()
	}
	detailContent := p.PeekDetailPaneStack().View()

	tabLineStyle := Style.Width(p.Width()).Border(lipgloss.NormalBorder(), false, false, true, false).Render(p.TabLine.View())
	statusLineStyle := Style.Width(p.Width()).Border(lipgloss.NormalBorder(), true, false, false, false).Render(p.StatusLine.View())

	mainHeight = p.Height() - lipgloss.Height(tabLineStyle) - lipgloss.Height(statusLineStyle)
	detailHeight = p.Height() - lipgloss.Height(tabLineStyle) - lipgloss.Height(statusLineStyle)

	mainStyled := Style.Height(mainHeight).Width(mainWidth).Border(lipgloss.NormalBorder(), false, true, false, false).Padding(0, 1).Render(mainContent)
	detailStyled := Style.Padding(0, 1).Render(detailContent)

	contentView := lipgloss.JoinHorizontal(lipgloss.Left, mainStyled, detailStyled)

	return lipgloss.JoinVertical(lipgloss.Top, tabLineStyle, contentView, statusLineStyle)
}

func (p *paneManager) AddPane(pane engine.ManagedPane) engine.PaneID {
	p.currentID++
	id := p.currentID
	pane.SetID(id)
	p.Panes[id] = pane

	p.Panes[id].Init()

	return id
}

func (p *paneManager) RemovePane(id engine.PaneID) { delete(p.Panes, id) }
