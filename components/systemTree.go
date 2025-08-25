package components

import (
	. "github.com/elitracy/planets/system"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

type SystemTree struct {
	Systems []System
	Tree    *widgets.Tree
	active  bool
}

func (st *SystemTree) SetActive(active bool) {
	st.active = active
	if active {
		st.Tree.BorderStyle = ui.NewStyle(ui.ColorGreen)
	} else {
		st.Tree.BorderStyle = ui.NewStyle(ui.ColorBlue)
	}

}
func (st *SystemTree) SetRect(x1, y1, x2, y2 int) {
	st.Tree.SetRect(x1, y1, x2, y2)
}

func (st *SystemTree) GetDrawable() ui.Drawable {
	return st.Tree
}

func (st *SystemTree) Init() {
	st.Tree = widgets.NewTree()

	nodes := []*widgets.TreeNode{}
	for _, s := range st.Systems {

		planetNodes := []*widgets.TreeNode{}

		for _, p := range s.Planets {
			node := &widgets.TreeNode{
				Value: nodeValue(p.Name),
				Nodes: nil,
			}
			planetNodes = append(planetNodes, node)
		}

		systemNode := &widgets.TreeNode{
			Value: nodeValue(s.Name),
			Nodes: planetNodes,
		}

		nodes = append(nodes, systemNode)
	}
	st.Tree.SetNodes(nodes)
}

func (st *SystemTree) Render() {
	st.Tree.TextStyle = ui.NewStyle(ui.ColorYellow)
	st.Tree.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)
	st.Tree.WrapText = false
}

func (st *SystemTree) Update() {
}

func (st *SystemTree) HandleKey(key string) {
	previousKey := ""
	switch key {
	case "j", "<Down>":
		st.Tree.ScrollDown()
	case "k", "<Up>":
		st.Tree.ScrollUp()
	case "<C-d>":
		st.Tree.ScrollHalfPageDown()
	case "<C-u>":
		st.Tree.ScrollHalfPageUp()
	case "<C-f>":
		st.Tree.ScrollPageDown()
	case "<C-b>":
		st.Tree.ScrollPageUp()
	case "g":
		if previousKey == "g" {
			st.Tree.ScrollTop()
		}
	case "<Home>":
		st.Tree.ScrollTop()
	case "<Enter>":
		st.Tree.ToggleExpand()
	case "G", "<End>":
		st.Tree.ScrollBottom()
	case "E":
		st.Tree.ExpandAll()
	case "C":
		st.Tree.CollapseAll()

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = key
		}
	}
}

func (st *SystemTree) DebugInfo() string {
	return st.Tree.GetRect().String()
}
