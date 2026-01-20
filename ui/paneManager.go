package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
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
	UITick          core.Tick
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

	return pm
}

func (p *paneManager) ActiveDetailPane() ManagedPane {
	if len(p.DetailPaneStack) == 0 {
		return NewErrorPane("No detail selected")
	}

	return p.DetailPaneStack[len(p.DetailPaneStack)-1]
}

func (p *paneManager) AddTab(pane ManagedPane) {
	p.TabLine.tabs = append(p.TabLine.tabs, pane)
}

func (p *paneManager) SetMainPane(pane ManagedPane) {
	p.MainPane = pane
}

func (p *paneManager) PushDetailPane(pane ManagedPane) {
	p.DetailPaneStack = append(p.DetailPaneStack, pane)
}

func (p *paneManager) PopDetailPane() {
	if len(p.DetailPaneStack) == 0 {
		return
	}

	p.DetailPaneStack = p.DetailPaneStack[:len(p.DetailPaneStack)-1]
}
func (p *paneManager) FlushDetailPane() {
	p.DetailPaneStack = nil
}

func (p *paneManager) Init() tea.Cmd {

	if p.MainPane == nil {
		p.SetMainPane(p.TabLine.tabs[0])
	}

	var cmds []tea.Cmd

	for i := range p.Panes {
		cmds = append(cmds, p.Panes[i].Init())
		// cmds = append(cmds, paneResizeCmd(i, p.Pane.width, p.Pane.height))
	}
	cmds = append(cmds, core.TickCmd(p.UITick))

	logging.Info("main dims: %vx%v", mainWidth, mainHeight)
	logging.Info("detail dims: %vx%v", detailWidth, detailHeight)
	cmds = append(cmds, paneResizeCmd(p.MainPane.ID(), mainWidth, mainHeight))
	cmds = append(cmds, paneResizeCmd(p.ActiveDetailPane().ID(), detailWidth, detailHeight))
	return tea.Batch(cmds...)
}

func (p *paneManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if p.MainPane == nil {
		p.MainPane = NewErrorPane("No content selected")
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd
	var model tea.Model

	switch msg := msg.(type) {
	case paneResizeMsg:
		for _, pane := range p.Panes {
			if pane.ID() == msg.paneID {
				model, _ = p.Panes[msg.paneID].Update(msg)
				p.Panes[msg.paneID] = model.(ManagedPane)
			}
		}
	case setMainFocusMsg:
		p.SetMainPane(p.Panes[msg.id])

		cmds = append(cmds, paneResizeCmd(p.MainPane.ID(), mainWidth, mainHeight))
		cmds = append(cmds, paneResizeCmd(p.ActiveDetailPane().ID(), detailWidth, detailHeight))
	case pushDetailStackMsg:
		p.PushDetailPane(p.Panes[msg.id])

		cmds = append(cmds, paneResizeCmd(p.MainPane.ID(), mainWidth, mainHeight))
		cmds = append(cmds, paneResizeCmd(p.ActiveDetailPane().ID(), detailWidth, detailHeight))
		logging.Info("pushed new detail")
		logging.Info("commands after push: %v", cmds)
	case popDetailStackMsg:
		p.PopDetailPane()

		cmds = append(cmds, paneResizeCmd(p.MainPane.ID(), mainWidth, mainHeight))
		cmds = append(cmds, paneResizeCmd(p.ActiveDetailPane().ID(), detailWidth, detailHeight))
	case flushDetailStackMsg:
		p.FlushDetailPane()

		cmds = append(cmds, paneResizeCmd(p.MainPane.ID(), mainWidth, mainHeight))
		cmds = append(cmds, paneResizeCmd(p.ActiveDetailPane().ID(), detailWidth, detailHeight))
	case tea.WindowSizeMsg:
		p.Pane.width = msg.Width
		p.Pane.height = msg.Height

		cmds = append(cmds, paneResizeCmd(p.MainPane.ID(), mainWidth, mainHeight))
		cmds = append(cmds, paneResizeCmd(p.ActiveDetailPane().ID(), detailWidth, detailHeight))
	case core.TickMsg:
		p.UITick++
		cmds = append(cmds, core.TickCmd(p.UITick))
	case tea.KeyMsg:

		switch msg.String() {
		case "tab", "shift+tab":
			_, cmd := p.TabLine.Update(msg)
			return p, cmd
		}

		if len(p.DetailPaneStack) == 0 {
			model, cmd = p.MainPane.Update(msg)
			p.MainPane = model.(ManagedPane)
		} else {
			detailPane := p.ActiveDetailPane()
			model, cmd = detailPane.Update(msg)
			detailPane = model.(ManagedPane)
		}

		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return p, tea.Batch(cmds...)
}

func (p *paneManager) View() string {
	mainContent := "No tab selected"
	if p.MainPane != nil {
		mainContent = p.MainPane.View()
	}
	detailContent := p.ActiveDetailPane().View()

	tablineStyle := consts.Style.Width(p.width).Border(lipgloss.NormalBorder(), false, false, true, false).Render(p.TabLine.View())

	mainHeight = p.height - lipgloss.Height(tablineStyle)
	detailHeight = p.height - lipgloss.Height(tablineStyle)

	mainStyled := consts.Style.Height(mainHeight).Width(mainWidth).Border(lipgloss.NormalBorder(), false, true, false, false).Padding(0, 1).Render(mainContent)
	detailStyled := consts.Style.Padding(0, 1).Render(detailContent)

	contentView := lipgloss.JoinHorizontal(lipgloss.Left, mainStyled, detailStyled)

	return lipgloss.JoinVertical(lipgloss.Top, tablineStyle, contentView)
}

func (p *paneManager) AddPane(pane ManagedPane) core.PaneID {
	p.currentID++

	pane.SetID(p.currentID)
	p.Panes[p.currentID] = pane

	logging.Info("%v: %v", p.currentID, pane.Title())
	p.Panes[p.currentID].Init()

	return p.currentID
}

func (p *paneManager) RemovePane(id core.PaneID) { delete(p.Panes, id) }
