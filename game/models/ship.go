package models

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
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
		Velocity:   engine.Velocity{X: config.SCOUT_VELOCITY, Y: config.SCOUT_VELOCITY, Z: config.SCOUT_VELOCITY},
	}

	return ship
}

type ShipManager struct {
	Ships     map[EntityID]*Ship
	currentID EntityID
}

func NewShipManager() *ShipManager {
	return &ShipManager{
		Ships: make(map[EntityID]*Ship),
	}
}

func (m ShipManager) GetShip(id EntityID) *Ship { return m.Ships[id] }

func (m *ShipManager) AddShip(ship *Ship) {
	ship.ID = m.GetNextID()
	m.Ships[ship.GetID()] = ship
}

func (m *ShipManager) GetNextID() EntityID {
	m.currentID++
	return m.currentID
}
