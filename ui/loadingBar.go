package ui

import (
	"math"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/state"
)

type LoadingBarPane struct {
	*Pane

	startTick core.Tick
	endTick   core.Tick
	progress  progress.Model
}

func NewLoadingBarPane(startTick, endTick core.Tick) *LoadingBarPane {
	return &LoadingBarPane{
		Pane:      &Pane{},
		progress:  progress.New(progress.WithDefaultGradient()),
		startTick: startTick,
		endTick:   endTick,
	}
}

func (p *LoadingBarPane) Init() tea.Cmd { return nil }

func (p *LoadingBarPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		p.width = msg.width
		p.height = msg.height
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popMainFocusCmd(p.Pane.id)
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	case core.UITickMsg:
		if state.State.CurrentTick >= p.endTick {
			cmd := p.progress.SetPercent(1.0)
			return p, cmd
		}

		if state.State.CurrentTick >= p.startTick {
			duration := float64(p.endTick - p.startTick)
			elapsed := float64(state.State.CurrentTick - p.startTick)
			percent := elapsed / duration
			cmd := p.progress.SetPercent(percent)
			return p, cmd
		}

		return p, nil

	case progress.FrameMsg:
		progressModel, cmd := p.progress.Update(msg)
		p.progress = progressModel.(progress.Model)
		return p, cmd
	}
	return p, nil
}

func (p *LoadingBarPane) View() string {
	filled := int(math.Ceil((float64(p.width) * p.progress.Percent())))
	return strings.Repeat("█", filled) + strings.Repeat("░", p.width-filled)
}
