package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
)

type setMainFocusMsg struct{ id core.PaneID }
type popMainFocusMsg struct{ id core.PaneID }

type pushDetailStackMsg struct{ id core.PaneID }
type popDetailStackMsg struct{ id core.PaneID }
type flushDetailStackMsg struct{}

type focusTabsMsg struct{ lastActiveID core.PaneID }

type paneResizeMsg struct {
	paneID core.PaneID
	width  int
	height int
}

func setMainFocusCmd(id core.PaneID) tea.Cmd { return func() tea.Msg { return setMainFocusMsg{id} } }
func popMainFocusCmd(id core.PaneID) tea.Cmd { return func() tea.Msg { return popMainFocusMsg{id} } }

func pushDetailStackCmd(id core.PaneID) tea.Cmd {
	return func() tea.Msg { return pushDetailStackMsg{id} }
}
func popDetailStackCmd(id core.PaneID) tea.Cmd {
	return func() tea.Msg { return popDetailStackMsg{id} }
}
func flushDetailStackCmd() tea.Cmd {
	return func() tea.Msg { return flushDetailStackMsg{} }
}

func paneResizeCmd(id core.PaneID, width, height int) tea.Cmd {
	return func() tea.Msg { return paneResizeMsg{paneID: id, width: width, height: height} }
}
