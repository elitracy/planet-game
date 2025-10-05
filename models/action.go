package models

import "github.com/elitracy/planets/models/constructions"

//go:generate stringer -type=ActionType
type ActionType int

const (
	BuildFarm ActionType = iota
	BuildMine
	BuildSolarGrid
	BuildColony
)

type Action struct {
	ID           int
	TargetEntity Entity
	Description  string
	Type         ActionType
	ExecuteTime  int
	Status       EventStatus
	*Order
}

func (a Action) GetID() int {
	return a.ID
}

func (a Action) GetStart() int {
	return a.ExecuteTime
}

func (a Action) GetDuration() int {
	return a.ExecuteTime - a.Order.ExecuteTime
}

func (a Action) GetStatus() EventStatus {
	return a.Status
}

func (a *Action) Execute() {
	a.Status = Executing

	if a.Type == BuildFarm {
		farm := constructions.CreateFarm(1)

		if planetEntity, ok := a.TargetEntity.(*Planet); ok {
			planetEntity.Constructions.Farms = append(planetEntity.Constructions.Farms, farm)
		}

		a.Status = Complete
	}
}
