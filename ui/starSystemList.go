package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/models"
	"github.com/elitracy/planets/game/orders"
)

var filteredSystems []*models.StarSystem

type StarSystemListPane struct {
	*engine.Pane

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
		Pane:      engine.NewPane(title, engine.NewKeyBindings()),
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
	p.GetKeys().
		Set(engine.Quit, "q").
		Set(engine.Select, "enter").
		Set(engine.Down, "j").
		Set(engine.Up, "k").
		Set(engine.Search, "/")
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
		if msg.paneID == p.Pane.ID() {
			p.SetSize(msg.width-2, msg.height)

			return p, nil
		}
	case engine.TickMsg:

		p.GetKeys().
			Unset(engine.Scout).
			Unset(engine.Select)

		system := p.systems[p.cursor]
		if !system.Scouted && !system.Colonized {
			p.GetKeys().Set(engine.Scout, "s")
		}

		if system.Scouted || system.Colonized {
			p.GetKeys().Set(engine.Select, "enter")
		}

	case tea.KeyMsg:
		if p.searching {
			switch msg.String() {
			case "esc":
				p.textInput.Blur()
				p.searching = false
				p.cursor = 0
				p.GetKeys().Unset(engine.Back)
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
		case p.GetKeys().Get(engine.Search):
			p.searching = true
			p.textInput.Focus()
			p.cursor = 0
			p.GetKeys().Set(engine.Back, "esc")

			return p, textinput.Blink
		case p.GetKeys().Get(engine.Scout):
			return p.handleScoutOrder()
		case p.GetKeys().Get(engine.Select):
			system := filteredSystems[p.cursor]
			systemInfoPane := NewSystemInfoPane(system.Name, system)
			paneID := PaneManager.AddPane(systemInfoPane)
			return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
		case p.GetKeys().Get(engine.Up):
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

		case p.GetKeys().Get(engine.Down):
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

		case p.GetKeys().Get(engine.Quit):
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
		if system.Colonized {
			row += " (colonized)"
		}

		if !system.Colonized && system.Scouted {
			row += " (scouted)"
		}

		if i == p.cursor {
			row = p.theme.FocusedStyle.Render(row)
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

	systemList = Style.Width(p.Width()).Padding(0, 1).Border(lipgloss.RoundedBorder(), true, false, false, false).Render(systemList)

	infoContainer := lipgloss.JoinVertical(lipgloss.Left, p.textInput.View(), systemList)

	content := lipgloss.JoinVertical(lipgloss.Left, infoContainer)
	return content
}

func (p *StarSystemListPane) handleScoutOrder() (tea.Model, tea.Cmd) {
	pane := CreateNewShipManagementPane(
		"Ship Management",
		&game.State.ShipManager,
		func(ship *models.Ship) {
			order := orders.NewScoutDestinationOrder(ship, models.Location{Position: p.systems[p.cursor].Location.Position, Entity: p.systems[p.cursor]}, game.State.CurrentTick+40)
			game.State.PushOrder(order)
		},
	)

	paneID := PaneManager.AddPane(pane)
	return p, tea.Sequence(pushDetailStackCmd(paneID), pushFocusStackCmd(paneID))
}
