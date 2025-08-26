package models

type StarSystem struct {
	Name    string
	Planets []*Planet
	Location
}
