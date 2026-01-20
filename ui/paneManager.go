package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
)

var mainWidthPercentage float32

var mainWidth int
var mainHeight int

var detailWidth int
var detailHeight int

type paneManager struct {
	*Pane

	TabLine         *TablinePane
	MainPane        ManagedPane
	DetailPaneStack []ManagedPane
	Panes           map[core.PaneID]ManagedPane
	currentID       core.PaneID
	focusStack      []core.PaneID

	UITick core.Tick
}

var PaneManager = NewPaneManager()

func NewPaneManager() *paneManager {

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		logging.Error("Failed to intialize Pane Manager: %v", err)
		return nil
	}

	pm := &paneManager{
		Panes:     make(map[core.PaneID]ManagedPane),
		currentID: 0,
		UITick:    0,
		TabLine:   NewTablinePane("Tabs", []ManagedPane{}),
		Pane: &Pane{
			width:  width,
			height: height,
		},
	}

	mainWidthPercentage = .4

	mainWidth = int(float32(pm.width) * mainWidthPercentage)
	detailWidth = int(float32(pm.width) * (1 - mainWidthPercentage))

	mainHeight = pm.height
	detailHeight = pm.height

	pm.Panes[-1] = NewErrorPane("No content.")

	return pm
}

func (p *paneManager) PushFocusStack(id core.PaneID) {
	p.focusStack = append(p.focusStack, id)
}

func (p *paneManager) PopFocusStack() {
	if len(p.focusStack) <= 0 {
		return
	}

	p.focusStack = p.focusStack[:len(p.focusStack)-1]
}

func (p *paneManager) PeekFocusStack() core.PaneID {
	if len(p.focusStack) <= 0 {
		return -1
	}

	return p.focusStack[len(p.focusStack)-1]
}

func (p *paneManager) AddTab(pane ManagedPane) {
	p.TabLine.tabs = append(p.TabLine.tabs, pane)
}

func (p *paneManager) SetMainPane(pane ManagedPane) {
	p.MainPane = pane
}

func (p *paneManager) PushDetailPaneStack(pane ManagedPane) {
	p.DetailPaneStack = append(p.DetailPaneStack, pane)
}

func (p *paneManager) PopDetailPaneStack() {
	if len(p.DetailPaneStack) <= 1 {
		return
	}

	pane := p.DetailPaneStack[len(p.DetailPaneStack)-1]
	p.DetailPaneStack = p.DetailPaneStack[:len(p.DetailPaneStack)-1]

	p.RemovePane(pane.ID())
}

func (p *paneManager) PeekDetailPaneStack() ManagedPane {
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
		p.SetMainPane(p.TabLine.tabs[0])
	}

	if p.TabLine != nil {
		cmds = append(cmds, p.TabLine.Init())
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
		core.TickCmd(p.UITick),
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
			p.Panes[msg.paneID] = model.(ManagedPane)
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
		p.Pane.width = msg.Width
		p.Pane.height = msg.Height
	case core.TickMsg:
		p.UITick++
		cmds := []tea.Cmd{core.TickCmd(p.UITick)}
		for id, pane := range p.Panes {
			model, cmd := pane.Update(msg)
			p.Panes[id] = model.(ManagedPane)

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
			p.Panes[p.PeekFocusStack()] = model.(ManagedPane)
			return p, cmd
		}
	default:
		var cmds []tea.Cmd
		for id, pane := range p.Panes {
			model, cmd := pane.Update(msg)
			p.Panes[id] = model.(ManagedPane)
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

	tablineStyle := Style.Width(p.width).Border(lipgloss.NormalBorder(), false, false, true, false).Render(p.TabLine.View())

	mainHeight = p.height - lipgloss.Height(tablineStyle)
	detailHeight = p.height - lipgloss.Height(tablineStyle)

	mainStyled := Style.Height(mainHeight).Width(mainWidth).Border(lipgloss.NormalBorder(), false, true, false, false).Padding(0, 1).Render(mainContent)
	detailStyled := Style.Padding(0, 1).Render(detailContent)

	contentView := lipgloss.JoinHorizontal(lipgloss.Left, mainStyled, detailStyled)

	return lipgloss.JoinVertical(lipgloss.Top, tablineStyle, contentView)
}

func (p *paneManager) AddPane(pane ManagedPane) core.PaneID {
	p.currentID++
	pane.SetID(p.currentID)
	p.Panes[p.currentID] = pane

	p.Panes[p.currentID].Init()

	return p.currentID
}

func (p *paneManager) RemovePane(id core.PaneID) { delete(p.Panes, id) }
