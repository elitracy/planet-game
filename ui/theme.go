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
	focusedStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("205")),
	blurredStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	cursorStyle:         lipgloss.NewStyle().Foreground(lipgloss.Color("205")),
	noStyle:             lipgloss.NewStyle(),
	helpStyle:           lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	cursorModeHelpStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("244")),
}
