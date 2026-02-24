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
	SetID(EntityID)
	GetName() string
	GetLocation() Location
	GetOrders() []task.Task
}

func (e CoreEntity) GetID() EntityID        { return e.ID }
func (e *CoreEntity) SetID(id EntityID)     { e.ID = id }
func (e CoreEntity) GetName() string        { return e.Name }
func (e CoreEntity) GetLocation() Location  { return e.Location }
func (e CoreEntity) GetOrders() []task.Task { return e.OrderQueue }

type EntityManager struct {
	Entities  map[EntityID]Entity
	currentID EntityID
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities: make(map[EntityID]Entity),
	}
}

func (em *EntityManager) Add(e Entity) {
	e.SetID(em.currentID)
	em.Entities[em.currentID] = e

	em.currentID++
}

func (em *EntityManager) Get(id EntityID) Entity {
	if entity, ok := em.Entities[id]; ok {
		return entity
	}

	return nil
}
