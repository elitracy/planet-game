package ui

type Pane interface {
	GetId() int
	SetId(int)
	GetTitle() string
}
