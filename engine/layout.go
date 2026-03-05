package engine

import (
	"slices"

	"github.com/charmbracelet/lipgloss"
)

type LayoutDirection int

const (
	LayoutHorizontal LayoutDirection = iota
	LayoutVertical
)

type LayoutNode struct {
	Pane      ManagedPane
	Direction LayoutDirection
	Ratio     float32
	Children  []*LayoutNode
}

func NewLayoutNode(pane ManagedPane, direction LayoutDirection, ratio float32) *LayoutNode {
	return &LayoutNode{
		Pane:      pane,
		Direction: direction,
		Ratio:     ratio,
	}
}

func (n *LayoutNode) isLeaf() bool {
	return n.Pane != nil
}

func (n *LayoutNode) FindChild(paneID PaneID) *LayoutNode {
	if n.isLeaf() && n.Pane.ID() == paneID {
		return n
	}

	for _, layout := range n.Children {
		p := layout.FindChild(paneID)
		if p != nil {
			return p
		}
	}

	return nil
}

func (n *LayoutNode) FindParent(paneID PaneID) (int, *LayoutNode) {
	for i, layout := range n.Children {
		if layout.isLeaf() && layout.Pane.ID() == paneID {
			return i, n
		} else {
			idx, l := layout.FindParent(paneID)
			if l != nil {
				return idx, l
			}
		}
	}

	return -1, nil
}

func (n *LayoutNode) RemoveChild(index int) {
	n.Children = append(n.Children[:index], n.Children[index+1:]...)
}

func (n *LayoutNode) Pop(paneID PaneID) *LayoutNode {
	if n.FindChild(paneID) == nil {
		Error("pane does not exist: %v", paneID)
		return nil
	}

	idx, parent := n.FindParent(paneID)

	if idx == -1 || parent == nil {
		return nil
	}

	child := parent.Children[idx]

	parent.RemoveChild(idx)

	if len(parent.Children) == 1 {
		parentRatio := parent.Ratio
		parent = parent.Children[0]
		parent.Children = nil
		parent.Ratio = parentRatio
	}

	for _, layout := range parent.Children {
		layout.Ratio /= (1 - child.Ratio)
	}

	return child
}

func (n *LayoutNode) Push(layout *LayoutNode, targetPaneID PaneID) {
	Info("PUSH: looking for target %v", targetPaneID)
	Info("PUSH: root isLeaf=%v, paneID=%v", n.isLeaf(), n.Pane)

	if n.FindChild(targetPaneID) == nil {
		Error("Target Pane ID does not exist: %v", targetPaneID)
		return
	}

	idx, parent := n.FindParent(targetPaneID)
	if idx == -1 {
		pane := n.Pane
		n.Pane = nil
		n.Direction = layout.Direction

		n.Children = append(n.Children,
			NewLayoutNode(pane, n.Direction, 1-layout.Ratio),
			layout,
		)

		return
	}

	child := parent.Children[idx]

	if parent.Direction == layout.Direction {
		parent.Children = slices.Insert(parent.Children, idx+1, layout)
		ratioScale := 1 - layout.Ratio
		for _, child := range parent.Children {
			child.Ratio *= ratioScale
		}
	} else {
		targetRatio := 1 - layout.Ratio
		targetPane := child.Pane
		child.Pane = nil
		child.Children = []*LayoutNode{NewLayoutNode(targetPane, layout.Direction, targetRatio), layout}
	}

}

func (n *LayoutNode) Render(width, height int) string {

	if n.isLeaf() {
		Info("RENDER leaf: %v", n.Pane.Title())
		n.Pane.SetSize(width, height)
		return n.Pane.View()
	}
	Info("RENDER branch: %v children, dir=%v", len(n.Children), n.Direction)

	var panes []string

	remaining := width
	if n.Direction == LayoutVertical {
		remaining = height
	}

	for i, child := range n.Children {
		var childWidth, childHeight = width, height

		switch n.Direction {
		case LayoutVertical:
			if i == len(n.Children)-1 {
				childHeight = remaining
			} else {
				childHeight = int(float32(height) * child.Ratio)
				remaining -= childHeight
			}
		case LayoutHorizontal:
			if i == len(n.Children)-1 {
				childWidth = remaining
			} else {
				childWidth = int(float32(width) * child.Ratio)
				remaining -= childWidth
			}

		}

		panes = append(panes, child.Render(childWidth, childHeight))
	}

	switch n.Direction {
	case LayoutVertical:
		return lipgloss.JoinVertical(lipgloss.Top, panes...)
	case LayoutHorizontal:
		return lipgloss.JoinHorizontal(lipgloss.Left, panes...)
	}

	return ""
}
