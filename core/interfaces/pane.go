package interfaces

type Pane interface {
	GetId() int
	SetId(int)
	GetTitle() string
	GetWidth() int
	GetHeight() int

	SetWidth(int)
	SetHeight(int)
}
