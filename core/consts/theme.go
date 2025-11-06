package consts

import "github.com/charmbracelet/lipgloss"

type UITheme struct {
	FocusedStyle        lipgloss.Style
	BlurredStyle        lipgloss.Style
	CursorStyle         lipgloss.Style
	NoStyle             lipgloss.Style
	HelpStyle           lipgloss.Style
	CursorModeHelpStyle lipgloss.Style
}

var Theme = UITheme{
	FocusedStyle:        Style.Foreground(lipgloss.Color("205")),
	BlurredStyle:        Style.Foreground(lipgloss.Color("240")),
	CursorStyle:         Style.Foreground(lipgloss.Color("205")),
	NoStyle:             Style,
	HelpStyle:           Style.Foreground(lipgloss.Color("240")),
	CursorModeHelpStyle: Style.Foreground(lipgloss.Color("244")),
}

var Style = lipgloss.NewStyle()
