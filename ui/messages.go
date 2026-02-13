package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/engine"
)

type setMainFocusMsg struct{ id engine.PaneID }

func setMainFocusCmd(id engine.PaneID) tea.Cmd { return func() tea.Msg { return setMainFocusMsg{id} } }

type popMainFocusMsg struct{ id engine.PaneID }

func popMainFocusCmd(id engine.PaneID) tea.Cmd { return func() tea.Msg { return popMainFocusMsg{id} } }

type pushDetailStackMsg struct{ id engine.PaneID }

func pushDetailStackCmd(id engine.PaneID) tea.Cmd {
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
	paneID engine.PaneID
	width  int
	height int
}

func paneResizeCmd(id engine.PaneID, width, height int) tea.Cmd {
	return func() tea.Msg { return paneResizeMsg{paneID: id, width: width, height: height} }
}

type pushFocusStackMsg struct{ id engine.PaneID }

func pushFocusStackCmd(id engine.PaneID) tea.Cmd {
	return func() tea.Msg { return pushFocusStackMsg{id} }
}

type popFocusStackMsg struct{}

func popFocusStackCmd() tea.Cmd { return func() tea.Msg { return popFocusStackMsg{} } }

type flushFocusStackMsg struct{}

func flushFocusStackCmd() tea.Cmd { return func() tea.Msg { return flushFocusStackMsg{} } }
