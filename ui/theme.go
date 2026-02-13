package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
)

// Blackbag color palette
const (
	ColorFg       = lipgloss.Color("#d0e8ff")
	ColorFgBright = lipgloss.Color("#ffffff")
	ColorFgDim    = lipgloss.Color("#a0c5e0")
	ColorFgMuted  = lipgloss.Color("#7099ba")
	ColorOrange   = lipgloss.Color("#ff8800")
	ColorRedLight = lipgloss.Color("#ff6b6b")
	ColorCyan     = lipgloss.Color("#66d9ef")
	ColorCyanDim  = lipgloss.Color("#4db8cc")
	ColorBlue     = lipgloss.Color("#78b9ff")
	ColorBlueDim  = lipgloss.Color("#5588cc")
	ColorTeal     = lipgloss.Color("#3eb489")
	ColorComment  = lipgloss.Color("#556677")
	ColorBorder   = lipgloss.Color("#5a6a7a")
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
	FocusedStyle:        Style.Foreground(ColorOrange),
	BlurredStyle:        Style.Foreground(ColorFgMuted),
	DimmedStyle:         Style.Foreground(ColorFgDim),
	CursorStyle:         Style.Foreground(ColorOrange),
	NoStyle:             Style.Foreground(ColorFg),
	HelpStyle:           Style.Foreground(ColorComment),
	CursorModeHelpStyle: Style.Foreground(ColorFgMuted),
}

func GetPaneTheme(pane engine.ManagedPane) UITheme {
	focused := PaneManager.PeekFocusStack() == pane.ID()

	var theme = Theme

	if !focused {
		theme.FocusedStyle = theme.BlurredStyle
	}

	return theme
}

var Style = lipgloss.NewStyle()
