package models

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
)

type ShipType int

const (
	Scout ShipType = iota
	Fighter
)

type Ship struct {
	*CoreEntity
	Velocity engine.Velocity
	ShipType ShipType
}

func CreateNewShip(name string, location Location, shipType ShipType) *Ship {
	ship := &Ship{
		CoreEntity: &CoreEntity{Name: name, Location: location},
		ShipType:   shipType,
		Velocity:   engine.Velocity{X: game.SCOUT_VELOCITY, Y: game.SCOUT_VELOCITY, Z: game.SCOUT_VELOCITY},
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
