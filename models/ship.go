package models

import (
	. "github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/interfaces"
)

type ShipType int

const (
	Scout ShipType = iota
	Fighter
)

type Ship struct {
	ID         int
	Name       string
	OrderQueue []interfaces.Event
	Position
	Velocity
	ShipType
}

func CreateNewShip(name string, position Position, shipType ShipType) *Ship {
	ship := &Ship{
		Name:     name,
		Position: position,
		ShipType: shipType,
		Velocity: Velocity{5, 5, 5},
	}

	return ship
}

func (s Ship) GetID() int                    { return s.ID }
func (s Ship) GetName() string               { return s.Name }
func (s Ship) GetPosition() Position         { return s.Position }
func (s Ship) GetOrders() []interfaces.Event { return s.OrderQueue }

type ShipManager struct {
	Ships     map[int]*Ship
	currentID int
}

func (m ShipManager) GetShip(id int) *Ship { return m.Ships[id] }

func (m *ShipManager) AddShip(ship *Ship) {
	ship.ID = m.GetNextID()
	m.Ships[ship.GetID()] = ship
}

func (m *ShipManager) GetNextID() int {
	m.currentID++
	return m.currentID
}
