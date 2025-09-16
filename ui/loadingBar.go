package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/logging"
	"strings"
)

const (
	padding  = 2
	maxWidth = 80
)

func NewLoadingBarPane(id int, title string) LoadingBarPane {
	return LoadingBarPane{progress: progress.New(progress.WithDefaultGradient())}
}

type LoadingBarPane struct {
	Pane
	id       int
	title    string
	progress progress.Model
}

func (p LoadingBarPane) GetId() int       { return p.id }
func (p LoadingBarPane) GetTitle() string { return p.title }

func (p LoadingBarPane) Init() tea.Cmd {
	return tick()
}

func (p LoadingBarPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return PopFocus(), nil
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	case tea.WindowSizeMsg:
		p.progress.Width = msg.Width - padding*2 - 4
		p.progress.Width = min(p.progress.Width, maxWidth)

		return p, nil

	case tickMsg:
		if p.progress.Percent() == 1.0 {
			return p, nil
		}
		logging.Log("Got tick msg", "LoadingBar")

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := p.progress.IncrPercent(0.05)
		return p, tea.Batch(tick(), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := p.progress.Update(msg)
		p.progress = progressModel.(progress.Model)
		return p, cmd

	default:
		return p, nil
	}
	return p, nil
}

func (p LoadingBarPane) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" + pad + p.progress.View() + "\n"
}
