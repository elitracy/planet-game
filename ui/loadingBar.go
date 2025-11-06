package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/state"
	. "github.com/elitracy/planets/models"
)

func NewLoadingBarPane(title string, startTick, endTick core.Tick) *LoadingBarPane {
	return &LoadingBarPane{
		title:     title,
		progress:  progress.New(progress.WithDefaultGradient()),
		startTick: startTick,
		endTick:   endTick,
	}
}

type LoadingBarPane struct {
	id     core.PaneID
	title  string
	width  int
	height int

	startTick core.Tick
	endTick   core.Tick
	progress  progress.Model
}

func (p LoadingBarPane) GetId() core.PaneID    { return p.id }
func (p *LoadingBarPane) SetId(id core.PaneID) { p.id = id }
func (p LoadingBarPane) GetTitle() string      { return p.title }
func (p LoadingBarPane) GetWidth() int         { return p.width }
func (p LoadingBarPane) GetHeight() int        { return p.height }
func (p *LoadingBarPane) SetWidth(w int)       { p.width = w }
func (p *LoadingBarPane) SetHeight(h int)      { p.height = h }

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

	case core.TickMsg:
		if p.progress.Percent() == 1.0 {
			return p, nil
		}

		if p.startTick <= State.Tick {
			duration := p.endTick - p.startTick

			if duration == 0 {
				cmd := p.progress.IncrPercent(1)
				return p, cmd
			}

			increment := (float64(TICKS_PER_SECOND) / float64(core.TICKS_PER_SECOND_UI)) / float64(duration)

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
