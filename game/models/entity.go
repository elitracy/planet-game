package models

import (
	"github.com/elitracy/planets/engine"
)

type CoreEntity struct {
	ID         int
	Name       string
	OrderQueue []engine.Event
	Location   Location
}

type Entity interface {
	GetID() int
	GetName() string
	GetLocation() Location
	GetOrders() []engine.Event
}

func (p CoreEntity) GetID() int                { return p.ID }
func (p CoreEntity) GetName() string           { return p.Name }
func (p CoreEntity) GetLocation() Location     { return p.Location }
func (p CoreEntity) GetOrders() []engine.Event { return p.OrderQueue }
