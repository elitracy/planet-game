package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type UITheme struct {
	FocusedStyle        lipgloss.Style
	BlurredStyle        lipgloss.Style
	DimmedStyle         lipgloss.Style
	CursorStyle         lipgloss.Style
	NoStyle             lipgloss.Style
	HelpStyle           lipgloss.Style
	CursorModeHelpStyle lipgloss.Style
}

var Theme = UITheme{
	FocusedStyle:        Style.Foreground(lipgloss.Color("205")),
	BlurredStyle:        Style.Foreground(lipgloss.Color("240")),
	DimmedStyle:         Style.Foreground(lipgloss.Color("248")),
	CursorStyle:         Style.Foreground(lipgloss.Color("205")),
	NoStyle:             Style,
	HelpStyle:           Style.Foreground(lipgloss.Color("240")),
	CursorModeHelpStyle: Style.Foreground(lipgloss.Color("244")),
}

func GetPaneTheme(pane ManagedPane) UITheme {
	focused := PaneManager.PeekFocusStack() == pane.ID()

	var theme = Theme

	if !focused {
		theme.FocusedStyle = theme.DimmedStyle
	}

	return theme
}

var Style = lipgloss.NewStyle()
