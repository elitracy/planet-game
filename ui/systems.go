package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
)

type SystemsPane struct {
	id     core.PaneID
	title  string
	width  int
	height int

	cursor    int
	searching bool
	textInput textinput.Model
	gamestate *models.GameState
}

func (p SystemsPane) GetId() core.PaneID    { return p.id }
func (p *SystemsPane) SetId(id core.PaneID) { p.id = id }
func (p SystemsPane) GetTitle() string      { return p.title }
func (p SystemsPane) GetWidth() int         { return p.width }
func (p SystemsPane) GetHeight() int        { return p.height }
func (p *SystemsPane) SetWidth(w int)       { p.width = w }
func (p *SystemsPane) SetHeight(h int)      { p.height = h }

func NewSystemsPane(title string, gamestate *models.GameState) *SystemsPane {
	ti := textinput.New()
	ti.Placeholder = "Search Star Systems \"/\""
	ti.Blur()
	ti.CharLimit = 156
	ti.Width = 36

	pane := &SystemsPane{
		title:     title,
		gamestate: gamestate,
		searching: false,
		textInput: ti,
	}

	return pane
}

func (p *SystemsPane) Init() tea.Cmd {
	return nil
}

func (p *SystemsPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.GetId() {
			p.width = msg.width - 2
			p.height = msg.height

			return p, nil
		}
	case core.TickMsg:
		return p, nil
	case tea.KeyMsg:
		if p.searching {
			switch msg.String() {
			case "esc":
				p.textInput.Blur()
				p.searching = false
			}

			var cmd tea.Cmd
			p.textInput, cmd = p.textInput.Update(msg)

			return p, cmd
		}

		switch msg.String() {
		case "/":
			p.searching = true
			p.textInput.Focus()
			return p, textinput.Blink
		case "enter":
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.gamestate.StarSystems)-1 {
				p.cursor++
			}
		case "esc":
			return p, popFocusCmd()
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *SystemsPane) View() string {

	systems := p.gamestate.StarSystems
	var filteredSystems []*models.StarSystem

	if len(p.textInput.Value()) == 0 {
		filteredSystems = systems
	} else {
		for _, s := range systems {
			if strings.Contains(strings.ToLower(s.Name), strings.ToLower(p.textInput.Value())) {
				filteredSystems = append(filteredSystems, s)
			}
		}
	}
	// search bar

	// system list (filtered)
	var systemRows []string
	for _, s := range filteredSystems {
		row := fmt.Sprintf("%v", s.Name)
		systemRows = append(systemRows, row)
	}

	// active system (main content)

	// system data

	systemList := lipgloss.JoinVertical(lipgloss.Left, systemRows...)
	systemList = consts.Style.Width(36).Padding(0, 1).Border(lipgloss.RoundedBorder(), true, false, false, false).Render(systemList)

	content := lipgloss.JoinVertical(lipgloss.Left, p.textInput.View(), systemList)

	return content
}
