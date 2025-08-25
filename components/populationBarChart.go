package components

import (
	"fmt"

	. "github.com/elitracy/planets/system"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type PopulationBarChart struct {
	BarChart *widgets.BarChart
	System   System
	active   bool
}

func (bc *PopulationBarChart) SetActive(active bool) {
	bc.active = active
	if active {
		bc.BarChart.BorderStyle = ui.NewStyle(ui.ColorGreen)
	} else {
		bc.BarChart.BorderStyle = ui.NewStyle(ui.ColorBlue)
	}
}

func (bc *PopulationBarChart) SetRect(x1, y1, x2, y2 int) {
	bc.BarChart.SetRect(x1, y1, x2, y2)
}

func (bc *PopulationBarChart) Init() {
	bc.BarChart = widgets.NewBarChart()
	bc.BarChart.Title = "Planet Populations"

	bc.BarChart.Data = []float64{}
	bc.BarChart.Labels = []string{}

	for _, p := range bc.System.Planets {
		bc.BarChart.Data = append(bc.BarChart.Data, float64(p.Population))
		bc.BarChart.Labels = append(bc.BarChart.Labels, p.Name)
	}
}

func (bc *PopulationBarChart) Render() {
	bc.BarChart.BarWidth = 10
	bc.BarChart.BarColors = []ui.Color{ui.ColorYellow, ui.ColorCyan}
	bc.BarChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	bc.BarChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	bc.BarChart.NumFormatter = func(f float64) string {
		return fmt.Sprintf("%.2f", f)
	}
}

func (bc *PopulationBarChart) Update() {
}

func (bc *PopulationBarChart) GetDrawable() ui.Drawable {
	return bc.BarChart
}

func (bc *PopulationBarChart) DebugInfo() string {
	return bc.BarChart.GetRect().String()
}

func (bc *PopulationBarChart) HandleKey(key string) {
}
