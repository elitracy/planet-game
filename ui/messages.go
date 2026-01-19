package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
)

type pushMainFocusMsg struct{ id core.PaneID }
type popMainFocusMsg struct{ id core.PaneID }

type pushDetailsFocusMsg struct{ id core.PaneID }
type popDetailsFocusMsg struct{ id core.PaneID }

type focusTabsMsg struct{ lastActiveID core.PaneID }

type paneResizeMsg struct {
	paneID core.PaneID
	width  int
	height int
}

func pushMainFocusCmd(id core.PaneID) tea.Cmd { return func() tea.Msg { return pushMainFocusMsg{id} } }
func popMainFocusCmd(id core.PaneID) tea.Cmd  { return func() tea.Msg { return popMainFocusMsg{id} } }

func pushDetailsFocusCmd(id core.PaneID) tea.Cmd {
	return func() tea.Msg { return pushDetailsFocusMsg{id} }
}
func popDetailsFocusCmd(id core.PaneID) tea.Cmd {
	return func() tea.Msg { return popDetailsFocusMsg{id} }
}

func focusTabsCmd(lastActiveID core.PaneID) tea.Cmd {
	return func() tea.Msg { return focusTabsMsg{lastActiveID: lastActiveID} }
}

func paneResizeCmd(id core.PaneID, width, height int) tea.Cmd {
	return func() tea.Msg { return paneResizeMsg{paneID: id, width: width, height: height} }
}
