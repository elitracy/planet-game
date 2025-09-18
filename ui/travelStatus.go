package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/models"
)

type TravelStatusPane struct {
	Pane
	id          int
	title       string
	p0          models.Position // origin
	p1          models.Position // destination
	origin      string
	destination string
	startTick   int
	endTick     int
	state       *models.GameState
}

func (p *TravelStatusPane) Init() tea.Cmd { return nil }
func (p *TravelStatusPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}

	}
	return p, nil
}

func (p *TravelStatusPane) View() string {
	totalTime := p.endTick - p.startTick

	totalWidth := 20

	ticksPassed := p.state.CurrentTick - p.startTick
	ticksRemaining := p.endTick - p.state.CurrentTick

	if ticksRemaining <= 0 {
		return "Arrived!"
	}

	output := "["

	currentStatus := float64(ticksPassed) / float64(totalTime) * float64(totalWidth)
	for range int(currentStatus) {
		output += "="
	}

	for range totalWidth - int(currentStatus) {
		output += " "
	}
	output += "]"

	return output

}

func (p TravelStatusPane) GetId() int       { return p.id }
func (p *TravelStatusPane) SetId(id int)    { p.id = id }
func (p TravelStatusPane) GetTitle() string { return p.title }

func NewTravelStatusPane(title string, id int, p0, p1 models.Position, origin, destination string, startTick, endTick int, state *models.GameState) *TravelStatusPane {

	return &TravelStatusPane{
		title:       title,
		p0:          p0,
		p1:          p1,
		origin:      origin,
		destination: destination,
		startTick:   startTick,
		endTick:     endTick,
		state:       state,
	}
}
