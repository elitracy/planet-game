package ui

import "github.com/charmbracelet/lipgloss"

type UITheme struct {
	focusedStyle        lipgloss.Style
	blurredStyle        lipgloss.Style
	cursorStyle         lipgloss.Style
	noStyle             lipgloss.Style
	helpStyle           lipgloss.Style
	cursorModeHelpStyle lipgloss.Style
}

var Theme = UITheme{
	focusedStyle:        Style.Foreground(lipgloss.Color("205")),
	blurredStyle:        Style.Foreground(lipgloss.Color("240")),
	cursorStyle:         Style.Foreground(lipgloss.Color("205")),
	noStyle:             Style,
	helpStyle:           Style.Foreground(lipgloss.Color("240")),
	cursorModeHelpStyle: Style.Foreground(lipgloss.Color("244")),
}

var Style = lipgloss.NewStyle()
