package interfaces

import "github.com/elitracy/planets/core"

type Pane interface {
	GetId() core.PaneID
	SetId(core.PaneID)
	GetTitle() string
	GetWidth() int
	GetHeight() int

	SetWidth(int)
	SetHeight(int)
}
