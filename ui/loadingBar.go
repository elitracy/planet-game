package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/state"
)

func NewLoadingBarPane(title string, startTick, endTick core.Tick) *LoadingBarPane {
	return &LoadingBarPane{
		Pane: &Pane{
			title: title,
		},
		progress:  progress.New(progress.WithDefaultGradient()),
		startTick: startTick,
		endTick:   endTick,
	}
}

type LoadingBarPane struct {
	*Pane

	startTick core.Tick
	endTick   core.Tick
	progress  progress.Model
}

func (p *LoadingBarPane) Init() tea.Cmd { return nil }

func (p *LoadingBarPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popMainFocusCmd(p.Pane.id)
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	case core.UITickMsg:
		if p.progress.Percent() == 1.0 {
			return p, nil
		}

		if p.startTick <= State.Tick {
			duration := p.endTick - p.startTick

			if duration == 0 {
				cmd := p.progress.IncrPercent(1)
				return p, cmd
			}

			increment := (float64(core.TICKS_PER_SECOND) / float64(core.TICKS_PER_SECOND_UI)) / float64(duration)

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
	return p.progress.View()
}
