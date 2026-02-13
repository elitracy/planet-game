package ui

import (
	"math"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/config"
)

type ProgressBarPane struct {
	*engine.Pane

	startTick engine.Tick
	endTick   engine.Tick
	progress  progress.Model
}

func NewProgressBarPane(startTick, endTick engine.Tick) *ProgressBarPane {
	return &ProgressBarPane{
		Pane:      engine.NewPane("Progress Bar", nil),
		progress:  progress.New(progress.WithDefaultGradient()),
		startTick: startTick,
		endTick:   endTick,
	}
}

func (p *ProgressBarPane) Init() tea.Cmd { return nil }

func (p *ProgressBarPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		p.SetSize(msg.width, msg.height)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popMainFocusCmd(p.Pane.ID())
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	case config.UITickMsg:
		if game.State.CurrentTick >= p.endTick {
			cmd := p.progress.SetPercent(1.0)
			return p, cmd
		}

		if game.State.CurrentTick >= p.startTick {
			duration := float64(p.endTick - p.startTick)
			elapsed := float64(game.State.CurrentTick - p.startTick)
			percent := elapsed / duration
			cmd := p.progress.SetPercent(percent)
			return p, cmd
		}

	case progress.FrameMsg:
		progressModel, cmd := p.progress.Update(msg)
		p.progress = progressModel.(progress.Model)
		return p, cmd
	}
	return p, nil
}

func (p *ProgressBarPane) View() string {
	filled := int(math.Ceil((float64(p.Width()) * p.progress.Percent())))
	return strings.Repeat("█", filled) + strings.Repeat("░", p.Width()-filled)
}
