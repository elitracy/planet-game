package models

type ActionType int

const (
	BuildFarm = iota
	BuildMine
	BuildSolarGrid
	BuildColony
)

type Action struct {
	ID           int
	TargetEntity *Entity
	Description  string
	Type         ActionType
	StartTime    int
	Duration     int
	Status       EventStatus
}

func (a Action) GetID() int {
	return a.ID
}

func (a Action) GetStart() int {
	return a.StartTime
}

func (a Action) GetDuration() int {
	return a.Duration
}

func (a Action) GetStatus() EventStatus {
	return a.Status
}
