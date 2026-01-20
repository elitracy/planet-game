package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/state"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/orders"
)

type CreateColonyPane struct {
	*Pane
	planet     *models.Planet
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	theme      UITheme
}

func (p *CreateColonyPane) Init() tea.Cmd {
	return textinput.Blink
}

func (p *CreateColonyPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case core.TickMsg:
		return p, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if p.focusIndex == len(p.inputs) {
				p.planet.ColonyName = p.inputs[0].Value()

				createColonyOrder := orders.NewCreateColonyOrder(p.planet, state.State.Tick)

				state.State.OrderScheduler.Push(createColonyOrder)
				return p, tea.Batch(popDetailStackCmd(), popFocusStackCmd())
			}

			if p.cursorMode > 0 {
				p.cursorMode--
			}

			return p, nil
		case "esc":
			if p.cursorMode < cursor.CursorHide {
				p.cursorMode++
			}

			if p.cursorMode >= cursor.CursorHide {
				return p, tea.Batch(popDetailStackCmd(), popFocusStackCmd())
			}

			cmds := make([]tea.Cmd, len(p.inputs))
			for i := range p.inputs {
				cmds[i] = p.inputs[i].Cursor.SetMode(p.cursorMode)
			}

			return p, tea.Batch(cmds...)
		case "j":
			if p.focusIndex < len(p.inputs) && p.cursorMode != cursor.CursorBlink {
				p.focusIndex++
				p.updateCursorStyles()
			}
		case "k":
			if p.focusIndex > 0 && p.cursorMode != cursor.CursorBlink {
				p.focusIndex--
				p.updateCursorStyles()
			}

		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	if p.cursorMode == cursor.CursorBlink {
		return p, p.updateInputs(msg)
	}
	return p, nil

}

func (p *CreateColonyPane) updateCursorStyles() (tea.Model, tea.Cmd) {

	cmds := make([]tea.Cmd, len(p.inputs))
	for i := 0; i <= len(p.inputs)-1; i++ {
		if i == p.focusIndex {
			// Set focused state
			cmds[i] = p.inputs[i].Focus()
			p.inputs[i].PromptStyle = p.theme.FocusedStyle
			p.inputs[i].TextStyle = p.theme.FocusedStyle
			continue
		}
		// Remove focused state
		p.inputs[i].Blur()
		p.inputs[i].PromptStyle = p.theme.NoStyle
		p.inputs[i].TextStyle = p.theme.NoStyle
	}

	return p, tea.Batch(cmds...)
}

func (p *CreateColonyPane) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(p.inputs))

	for i := range p.inputs {
		p.inputs[i], cmds[i] = p.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (p *CreateColonyPane) View() string {
	p.theme = GetPaneTheme(p)

	var (
		focusedButton = p.theme.FocusedStyle.Render("[ Submit ]")
		blurredButton = fmt.Sprintf("[ %s ]", p.theme.BlurredStyle.Render("Submit"))
	)

	var b strings.Builder

	for i := range p.inputs {
		b.WriteString(p.inputs[i].View())
		if i < len(p.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if p.focusIndex == len(p.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(p.theme.HelpStyle.Render("cursor mode is "))
	b.WriteString(p.theme.CursorModeHelpStyle.Render(p.cursorMode.String()))
	b.WriteString(p.theme.HelpStyle.Render(" (<esc> to change style)"))

	return b.String()
}

func NewCreateColonyPane(title string, planet *models.Planet) *CreateColonyPane {

	p := &CreateColonyPane{
		inputs: make([]textinput.Model, 4),
		Pane: &Pane{
			title: title,
		},
		planet:     planet,
		cursorMode: cursor.CursorStatic,
	}

	var t textinput.Model
	for i := range p.inputs {
		t = textinput.New()
		t.Cursor.Style = p.theme.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Colony Name"
			t.CharLimit = 64
			t.Width = 20
			t.Focus()
			t.PromptStyle = p.theme.FocusedStyle
			t.TextStyle = p.theme.FocusedStyle
		case 1:
			t.Placeholder = "Food"
			t.CharLimit = 64
			t.Width = 20
		case 2:
			t.Placeholder = "Minerals"
			t.CharLimit = 64
			t.Width = 20
		case 3:
			t.Placeholder = "Energy"
			t.CharLimit = 64
			t.Width = 20
		}

		p.inputs[i] = t
	}

	return p
}
