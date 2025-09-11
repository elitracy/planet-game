package ui

type BasePane struct {
	id    int
	title string
}

func (p BasePane) GetId() int {
	return p.id
}

func (p BasePane) GetTitle() string {
	return p.title
}
