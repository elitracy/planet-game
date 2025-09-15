package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/elitracy/planets/models"
)

type PlayerInfoPane struct {
	Pane
	id        int
	title     string
	gamestate *models.GameState

	selected    int
	cursor      int
	max_choices int
	prev_key    string
}

func NewPlayerInfoPane(text string, id int, gs *models.GameState) *PlayerInfoPane {
	max_choices := 0
	for _, s := range gs.StarSystems {
		max_choices++
		for range s.Planets {
			max_choices++
		}
	}
	return &PlayerInfoPane{title: text, id: id, gamestate: gs, max_choices: max_choices, cursor: 1}
}

func (p *PlayerInfoPane) Init() tea.Cmd { return tick(p.id) }
func (p *PlayerInfoPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case TickMsg:
		if msg.id == p.id {
			return p, tick(p.id)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 1 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < p.max_choices {
				p.cursor++
			}

		case "G":
			p.cursor = p.max_choices
		case "g":
			if p.prev_key == "g" {
				p.cursor = 1
				p.prev_key = ""
			}
		case " ":
			p.selected = p.cursor
		case "esc":
			return PopFocus(), nil
		case "ctrl+c", "q":
			return p, tea.Quit
		}
		p.prev_key = msg.String()
	}

	return p, nil
}

func (p *PlayerInfoPane) View() string {

	title := p.title + "\n"
	title += fmt.Sprintf("Position: %v\n", p.gamestate.Player.Position)

	content := "Destinations: \n"

	t := tree.New()
	t.Root("Destinations")

	current_row := 0
	for _, system := range p.gamestate.StarSystems {
		cursor := ""
		system_branch := tree.New()

		current_row++
		if p.cursor == current_row {
			cursor = ">"
		}

		system_branch.Root(fmt.Sprintf("%s %s", cursor, system.Name))

		for _, planet := range system.Planets {
			current_row++
			cursor = ""

			if p.cursor == current_row {
				cursor = ">"
			}

			checked := " "
			if p.selected == current_row {
				checked = "x"
			}
			system_branch.Child(fmt.Sprintf("%s [%s] %s", cursor, checked, planet.Name))
		}

		t.Child(system_branch)
	}
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("63")).MarginRight(1)
	rootStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("35"))
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	t.Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(enumeratorStyle).
		RootStyle(rootStyle).
		ItemStyle(itemStyle)

	content += t.String()
	content += "\n"

	title = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).Render(title)

	buttonContainer := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Align(lipgloss.Center).Padding(1, 2)

	confirmButton := buttonContainer.Render("Start Travel")
	cancelButton := buttonContainer.Render("Cancel")

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, confirmButton, cancelButton)

	block := lipgloss.Place(10, 2, lipgloss.Center, lipgloss.Top, title)
	block += lipgloss.Place(10, 20, lipgloss.Left, lipgloss.Center, content)
	block += lipgloss.Place(10, 10, lipgloss.Center, lipgloss.Bottom, buttons)

	return block
}

func (p PlayerInfoPane) GetId() int { return p.id }
