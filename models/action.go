package models

type ActionType int

const (
	BuildFarm = iota
	BuildMine
	BuildSolarGrid
	BuildColony
)

type Action struct {
	ID          int
	Type        ActionType
	Duration    int
	Progress    int
	Status      Status
	Description string
}

func (a *Action) Execute() Status {
	a.Status = Executing
	return a.Status
}
