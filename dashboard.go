package main

import (
	ui "github.com/gizak/termui/v3"
)

type Widget interface {
	Init()
	Update()
	Render()
	SetActive(active bool)
	SetRect(x1, y1, x2, y2 int)
	DebugInfo() string
	GetDrawable() ui.Drawable
	HandleKey(key string)
}

type Dashboard struct {
	Components [][]Widget
	ActiveRow  int
	ActiveCol  int
}

func (d *Dashboard) Render() {

	var drawables []ui.Drawable
	for row := range d.Components {
		for _, c := range d.Components[row] {
			c.Render()
			drawables = append(drawables, c.GetDrawable())
		}
	}

	ui.Render(drawables...)
}

func (d *Dashboard) SetActiveWidget(row, col int) {
	for r, line := range d.Components {
		for c, comp := range line {
			comp.SetActive(r == row && c == col)
		}
	}
}

func (d *Dashboard) GetActiveWidget() Widget {
	return d.Components[d.ActiveRow][d.ActiveCol]

}

func (d *Dashboard) MoveFocus(dRow, dCol int) {
	nextRow := d.ActiveRow + dRow
	nextCol := d.ActiveCol + dCol

	if nextRow < 0 || nextRow >= len(d.Components) {
		return
	}

	if nextCol < 0 || nextCol >= len(d.Components[nextRow]) {
		return
	}

	d.ActiveRow = nextRow
	d.ActiveCol = nextCol

	d.SetActiveWidget(d.ActiveRow, d.ActiveCol)
}

func (d *Dashboard) SetRects() {
	termWidth, termHeight := ui.TerminalDimensions()

	rows := len(d.Components)
	var rowCols []int
	for r := range len(d.Components) {
		rowCols = append(rowCols, len(d.Components[r]))
	}

	cellHeight := termHeight / rows

	for r, row := range d.Components {
		cellWidth := termWidth / rowCols[r]
		for c, comp := range row {
			x1 := c * cellWidth
			y1 := r * cellHeight
			x2 := x1 + cellWidth
			y2 := y1 + cellHeight

			comp.SetRect(x1, y1, x2, y2)
		}
	}

}

func (d *Dashboard) HandleKey(key string) {
	switch key {
	case "H":
		d.MoveFocus(0, -1)
		return
	case "L":
		d.MoveFocus(0, 1)
		return
	case "K":
		d.MoveFocus(-1, 0)
		return
	case "J":
		d.MoveFocus(1, 0)
		return
	case "<Resize>":
		d.SetRects()
		d.Render()
		return
	}

	d.GetActiveWidget().HandleKey(key)
	d.GetActiveWidget().Update()
}
