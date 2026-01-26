package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events/orders"
	"github.com/elitracy/planets/state"
)

var filteredSystems []*models.StarSystem

type StarSystemListPane struct {
	*Pane

	cursor    int
	searching bool
	textInput textinput.Model
	systems   []*models.StarSystem
	theme     UITheme
}

func NewStarSystemListPane(title string, systems []*models.StarSystem) *StarSystemListPane {
	ti := textinput.New()
	ti.Placeholder = "Search Star Systems \"/\""
	ti.Blur()
	ti.CharLimit = 156
	ti.Width = 36

	pane := &StarSystemListPane{
		Pane: &Pane{
			title: title,
			keys:  "Select: enter | Down: j | Up: k",
		},
		systems:   systems,
		searching: false,
		textInput: ti,
	}

	return pane
}

func (p *StarSystemListPane) filterSystems() {
	if len(p.textInput.Value()) == 0 {
		filteredSystems = p.systems
	} else {
		filteredSystems = []*models.StarSystem{}
		for _, s := range p.systems {
			if strings.Contains(strings.ToLower(s.Name), strings.ToLower(p.textInput.Value())) {
				filteredSystems = append(filteredSystems, s)
			}
		}
	}

	if p.cursor >= len(filteredSystems) {
		p.cursor = max(0, len(filteredSystems)-1)
	}
}

func (p *StarSystemListPane) Init() tea.Cmd {
	filteredSystems = p.systems

	if len(filteredSystems) == 0 {
		return nil
	}

	system := filteredSystems[p.cursor]
	systemInfoPane := NewSystemInfoPane(system.Name, system)
	paneID := PaneManager.AddPane(systemInfoPane)

	return pushDetailStackCmd(paneID)
}

func (p *StarSystemListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case paneResizeMsg:
		if msg.paneID == p.Pane.id {
			p.Pane.width = msg.width - 2
			p.Pane.height = msg.height

			return p, nil
		}
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

			p.filterSystems()

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

		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p *StarSystemListPane) View() string {
	p.theme = GetPaneTheme(p)

	var systemRows []string
	for i, system := range filteredSystems {
		row := fmt.Sprintf("%v", system.Name)
		if !system.Colonized {
			row += " (uncolonized)"

			if system.Scouted {
				row += " (scouted)"
			}
		}

		if i == p.cursor {
			if system.Colonized {
				row = p.theme.FocusedStyle.Render(row)
			}
		}

		if i != p.cursor {
			row = p.theme.DimmedStyle.Render(row)
		}
		systemRows = append(systemRows, row)

	}

	systemList := ""
	if len(systemRows) > 0 {
		systemList = lipgloss.JoinVertical(lipgloss.Left, systemRows...)
	}

	systemList = Style.Width(p.width).Padding(0, 1).Border(lipgloss.RoundedBorder(), true, false, false, false).Render(systemList)

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, p.textInput.View(), systemList)

	content := lipgloss.JoinVertical(lipgloss.Left, infoContainer)
	return content
}

func (p *StarSystemListPane) handleScoutOrder() (tea.Model, tea.Cmd) {
	pane := CreateNewShipManagementPane(
		"Ship Management",
		&state.State.ShipManager,
		func(ship *models.Ship) {
			order := orders.NewScoutDestinationOrder(ship, models.Destination{Position: p.systems[p.cursor].Position, Entity: p.systems[p.cursor]}, state.State.CurrentTick+40)
			state.State.OrderScheduler.Push(order)
		},
	)

	paneID := PaneManager.AddPane(pane)
	return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
}
