package models

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models/events"
)

type CoreEntity struct {
	ID         int
	Name       string
	OrderQueue []events.Event
	Position   core.Position
}

type Entity interface {
	GetID() int
	GetName() string
	GetPosition() core.Position
	GetOrders() []events.Event
}

func (p CoreEntity) GetID() int                 { return p.ID }
func (p CoreEntity) GetName() string            { return p.Name }
func (p CoreEntity) GetPosition() core.Position { return p.Position }
func (p CoreEntity) GetOrders() []events.Event  { return p.OrderQueue }
