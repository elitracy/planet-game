package models

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
)

type ShipType int

const (
	Scout ShipType = iota
	Fighter
)

type Ship struct {
	*CoreEntity
	Velocity core.Velocity
	ShipType ShipType
}

func CreateNewShip(name string, location Location, shipType ShipType) *Ship {
	ship := &Ship{
		CoreEntity: &CoreEntity{Name: name, Location: location},
		ShipType:   shipType,
		Velocity:   core.Velocity{X: consts.SCOUT_VELOCITY, Y: consts.SCOUT_VELOCITY, Z: consts.SCOUT_VELOCITY},
	}

	return ship
}

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
