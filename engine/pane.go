package engine

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Pane struct {
	id     PaneID
	title  string
	width  int
	height int
	keys   *KeyBindings
}

func (p Pane) String() string {
	return fmt.Sprintf("[%v] %v (%vx%v)", p.id, p.title, p.width, p.height)

}

func (p *Pane) ID() PaneID                { return p.id }
func (p *Pane) Width() int                { return p.width }
func (p *Pane) Height() int               { return p.height }
func (p *Pane) SetID(id PaneID)           { p.id = id }
func (p *Pane) Title() string             { return p.title }
func (p *Pane) Size() (int, int)          { return p.width, p.height }
func (p *Pane) SetSize(w, h int)          { p.width, p.height = w, h }
func (p *Pane) GetKeys() *KeyBindings     { return p.keys }
func (p *Pane) SetKeys(keys *KeyBindings) { p.keys = keys }

func NewPane(title string, keys *KeyBindings) *Pane {
	return &Pane{title: title, keys: keys}
}

type ManagedPane interface {
	tea.Model
	ID() PaneID
	Height() int
	Width() int
	SetID(PaneID)
	Title() string
	Size() (int, int)
	SetSize(int, int)
	GetKeys() *KeyBindings
	SetKeys(keys *KeyBindings)
}
