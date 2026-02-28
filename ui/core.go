package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/engine"
)

type setLayoutMsg struct {
	layout *engine.LayoutNode
}

func setLayoutCmd(layout *engine.LayoutNode) tea.Cmd {
	return func() tea.Msg { return setLayoutMsg{layout} }
}

type pushLayoutMsg struct {
	layout       *engine.LayoutNode
	targetPaneID engine.PaneID
}

func pushLayoutCmd(layout *engine.LayoutNode, targetPaneID engine.PaneID) tea.Cmd {
	layout.Pane.Init()
	return func() tea.Msg { return pushLayoutMsg{layout, targetPaneID} }
}

type popLayoutMsg struct{}

func popLayoutCmd() tea.Cmd {
	return func() tea.Msg { return popLayoutMsg{} }
}

type pushFocusStackMsg struct{ id engine.PaneID }

func pushFocusStackCmd(id engine.PaneID) tea.Cmd {
	return func() tea.Msg { return pushFocusStackMsg{id} }
}

type popFocusStackMsg struct{}

func popFocusStackCmd() tea.Cmd { return func() tea.Msg { return popFocusStackMsg{} } }

type flushFocusStackMsg struct{}

func flushFocusStackCmd() tea.Cmd { return func() tea.Msg { return flushFocusStackMsg{} } }
