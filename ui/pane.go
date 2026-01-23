package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elitracy/planets/core"
)

type Pane struct {
	id     core.PaneID
	title  string
	width  int
	height int
}

func (p Pane) String() string {
	return fmt.Sprintf("[%v] %v (%vx%v)", p.id, p.title, p.width, p.height)

}

func (p *Pane) ID() core.PaneID      { return p.id }
func (p *Pane) SetID(id core.PaneID) { p.id = id }
func (p *Pane) Title() string        { return p.title }
func (p *Pane) Size() (int, int)     { return p.width, p.height }
func (p *Pane) SetSize(w, h int)     { p.width, p.height = w, h }

type ManagedPane interface {
	tea.Model
	ID() core.PaneID
	SetID(core.PaneID)
	Title() string
	Size() (int, int)
	SetSize(int, int)
}
