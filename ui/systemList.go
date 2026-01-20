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

var filteredSystems []*models.StarSystem

type SystemsPane struct {
	*Pane

	cursor    int
	searching bool
	textInput textinput.Model
	gamestate *models.GameState
}

func NewSystemsPane(title string, gamestate *models.GameState) *SystemsPane {
	ti := textinput.New()
	ti.Placeholder = "Search Star Systems \"/\""
	ti.Blur()
	ti.CharLimit = 156
	ti.Width = 36

	pane := &SystemsPane{
		Pane: &Pane{
			title: title,
		},
		gamestate: gamestate,
		searching: false,
		textInput: ti,
	}

	return pane
}

func (p *SystemsPane) filteredSystems() {
	systems := p.gamestate.StarSystems
	if len(p.textInput.Value()) == 0 {
		filteredSystems = systems
	} else {
		filteredSystems = []*models.StarSystem{}
		for _, s := range systems {
			if strings.Contains(strings.ToLower(s.Name), strings.ToLower(p.textInput.Value())) {
				filteredSystems = append(filteredSystems, s)
			}
		}
	}

	if p.cursor >= len(filteredSystems) {
		p.cursor = max(0, len(filteredSystems)-1)
	}
}

func (p *SystemsPane) Init() tea.Cmd {
	filteredSystems = p.gamestate.StarSystems

	system := filteredSystems[p.cursor]
	systemInfoPane := NewSystemInfoPane(system.Name, system)
	paneID := PaneManager.AddPane(systemInfoPane)

	return pushDetailStackCmd(paneID)
}

func (p *SystemsPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.Pane.id {
			p.Pane.width = msg.width - 2
			p.Pane.height = msg.height

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
				p.cursor = 0
			}

			var cmd tea.Cmd
			var cmds []tea.Cmd
			p.textInput, cmd = p.textInput.Update(msg)

			cmds = append(cmds, cmd)

			p.filteredSystems()

			if len(filteredSystems) > 0 {
				system := filteredSystems[p.cursor]
				systemInfoPane := NewSystemInfoPane(system.Name, system)
				paneID := PaneManager.AddPane(systemInfoPane)
				cmds = append(cmds, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID)))
			}

			return p, tea.Sequence(cmds...)
		}

		switch msg.String() {
		case "/", "i", "a":
			p.searching = true
			p.textInput.Focus()
			p.cursor = 0

			return p, textinput.Blink
		case "enter":
			system := filteredSystems[p.cursor]
			systemInfoPane := NewSystemInfoPane(system.Name, system)
			paneID := PaneManager.AddPane(systemInfoPane)
			return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			} else {
				p.cursor = len(filteredSystems) - 1
			}

			if len(filteredSystems) == 0 {
				return p, nil
			}

			system := filteredSystems[p.cursor]
			systemInfoPane := NewSystemInfoPane(system.Name, system)
			paneID := PaneManager.AddPane(systemInfoPane)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case "down", "j":
			if p.cursor < len(filteredSystems)-1 {
				p.cursor++
			} else {
				p.cursor = 0
			}

			if len(filteredSystems) == 0 {
				return p, nil
			}
			system := filteredSystems[p.cursor]
			systemInfoPane := NewSystemInfoPane(system.Name, system)
			paneID := PaneManager.AddPane(systemInfoPane)
			return p, tea.Sequence(popDetailStackCmd(), pushDetailStackCmd(paneID))
		case "esc":
			return p, tea.Sequence(popDetailStackCmd(), popFocusStackCmd())
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *SystemsPane) View() string {
	var systemRows []string
	for i, s := range filteredSystems {
		row := fmt.Sprintf("%v", s.Name)
		if i == p.cursor {
			row = consts.Theme.FocusedStyle.Render(row)

		}
		systemRows = append(systemRows, row)

	}
	systemList := ""
	if len(systemRows) > 0 {
		systemList = lipgloss.JoinVertical(lipgloss.Left, systemRows...)
	}

	systemList = consts.Style.Width(36).Padding(0, 1).Border(lipgloss.RoundedBorder(), true, false, false, false).Render(systemList)

	content := lipgloss.JoinVertical(lipgloss.Left, p.textInput.View(), systemList)

	return content
}
