package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
)

type setMainFocusMsg struct{ id core.PaneID }

func setMainFocusCmd(id core.PaneID) tea.Cmd { return func() tea.Msg { return setMainFocusMsg{id} } }

type popMainFocusMsg struct{ id core.PaneID }

func popMainFocusCmd(id core.PaneID) tea.Cmd { return func() tea.Msg { return popMainFocusMsg{id} } }

type pushDetailStackMsg struct{ id core.PaneID }

func pushDetailStackCmd(id core.PaneID) tea.Cmd {
	return func() tea.Msg { return pushDetailStackMsg{id} }
}

type popDetailStackMsg struct{}

func popDetailStackCmd() tea.Cmd {
	return func() tea.Msg { return popDetailStackMsg{} }
}

type flushDetailStackMsg struct{}

func flushDetailStackCmd() tea.Cmd {
	return func() tea.Msg { return flushDetailStackMsg{} }
}

type paneResizeMsg struct {
	paneID core.PaneID
	width  int
	height int
}

func paneResizeCmd(id core.PaneID, width, height int) tea.Cmd {
	return func() tea.Msg { return paneResizeMsg{paneID: id, width: width, height: height} }
}

type pushFocusStackMsg struct{ id core.PaneID }

func pushFocusStackCmd(id core.PaneID) tea.Cmd {
	return func() tea.Msg { return pushFocusStackMsg{id} }
}

type popFocusStackMsg struct{}

func popFocusStackCmd() tea.Cmd { return func() tea.Msg { return popFocusStackMsg{} } }

type flushFocusStackMsg struct{}

func flushFocusStackCmd() tea.Cmd { return func() tea.Msg { return flushFocusStackMsg{} } }
