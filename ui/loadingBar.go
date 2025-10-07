package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
)

const (
	padding  = 2
	maxWidth = 80
)

func NewLoadingBarPane(title string, startTick, endTick int) *LoadingBarPane {
	return &LoadingBarPane{
		title:     title,
		progress:  progress.New(progress.WithDefaultGradient()),
		startTick: startTick,
		endTick:   endTick,
	}
}

type LoadingBarPane struct {
	Pane
	id        int
	title     string
	startTick int
	endTick   int
	progress  progress.Model
}

func (p LoadingBarPane) GetId() int       { return p.id }
func (p *LoadingBarPane) SetId(id int)    { p.id = id }
func (p LoadingBarPane) GetTitle() string { return p.title }

func (p *LoadingBarPane) Init() tea.Cmd { return nil }

func (p *LoadingBarPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popFocusCmd()
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

		if p.startTick >= GameStateGlobal.CurrentTick {
			duration := p.endTick - p.startTick

			if duration == 0 {
				cmd := p.progress.IncrPercent(1)
				return p, cmd
			}

			increment := 1.0 / float64(duration)

			logging.Info("Start Tick: %d", p.startTick)
			logging.Info("Game Tick: %d", GameStateGlobal.CurrentTick)
			logging.Info("increment: %.2f", increment)

			cmd := p.progress.IncrPercent(increment)
			return p, cmd
		}

		return p, nil

	case progress.FrameMsg:
		progressModel, cmd := p.progress.Update(msg)
		p.progress = progressModel.(progress.Model)
		return p, cmd

	default:
		return p, nil
	}
	return p, nil
}

func (p *LoadingBarPane) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" + pad + p.progress.View() + "\n"
}
