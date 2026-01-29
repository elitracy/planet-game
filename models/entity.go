package models

import (
	"github.com/elitracy/planets/models/events"
)

type CoreEntity struct {
	ID         int
	Name       string
	OrderQueue []events.Event
	Location   Location
}

type Entity interface {
	GetID() int
	GetName() string
	GetLocation() Location
	GetOrders() []events.Event
}

func (p CoreEntity) GetID() int                { return p.ID }
func (p CoreEntity) GetName() string           { return p.Name }
func (p CoreEntity) GetLocation() Location     { return p.Location }
func (p CoreEntity) GetOrders() []events.Event { return p.OrderQueue }
