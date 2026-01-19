package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
)

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

func (p *paneManager) PopDetailPane(pane ManagedPane) {
	if len(p.DetailPaneStack) == 0 {
		return
	}

	p.DetailPaneStack = p.DetailPaneStack[:len(p.DetailPaneStack)-1]
}

func (p *paneManager) Init() tea.Cmd {

	if p.MainPane == nil {
		p.SetMainPane(p.TabLine.tabs[0])
	}

	var cmds []tea.Cmd

	for i := range p.Panes {
		cmds = append(cmds, p.Panes[i].Init())
		cmds = append(cmds, paneResizeCmd(i, p.Pane.width, p.Pane.height))
	}
	cmds = append(cmds, core.TickCmd(p.UITick))

	return tea.Batch(cmds...)
}

func (p *paneManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if p.MainPane == nil {
		p.MainPane = NewErrorPane("No content selected")
	}

	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case pushMainFocusMsg:
		p.SetMainPane(p.Panes[msg.id])
	case pushDetailsFocusMsg:
		p.PushDetailPane(p.Panes[msg.id])
	case popDetailsFocusMsg:
		p.PushDetailPane(NewErrorPane("No details selected"))
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

			switch msg.String() {
			case "h", "l", "left", "right":
				_, cmd = p.TabLine.Update(msg)
				return p, cmd
			}

			if p.Panes[id].ID() == p.MainPane.ID() || p.Panes[id].ID() == p.ActiveDetailPane().ID() {
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

func (p *paneManager) View() string {
	mainContent := "No tab selected"
	if p.MainPane != nil {
		mainContent = p.MainPane.View()
	}

	detailContent := p.ActiveDetailPane().View()

	contentView := lipgloss.JoinHorizontal(lipgloss.Left, mainContent, detailContent)

	return lipgloss.JoinVertical(lipgloss.Top, p.TabLine.View(), contentView)
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
