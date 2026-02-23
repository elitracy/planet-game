package models

import (
	"github.com/elitracy/planets/engine/task"
)

type EntityID int

type CoreEntity struct {
	ID         EntityID
	Name       string
	OrderQueue []task.Task
	Location   Location
}

type Entity interface {
	GetID() EntityID
	GetName() string
	GetLocation() Location
	GetOrders() []task.Task
}

func (p CoreEntity) GetID() EntityID        { return p.ID }
func (p CoreEntity) GetName() string        { return p.Name }
func (p CoreEntity) GetLocation() Location  { return p.Location }
func (p CoreEntity) GetOrders() []task.Task { return p.OrderQueue }
