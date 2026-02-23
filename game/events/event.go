package events

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/models"
)

type EventSeverity int

const (
	Low EventSeverity = iota
	Moderate
	High
)

type EventID int

type EventChance struct {
	Probability float64
	New         func(models.Entity, engine.Tick) *Event
}

type Event struct {
	ID          EventID
	Name        string
	Description string
	Severity    EventSeverity
	Target      models.Entity
	Effect      func()
	Expire      func()
	Resolved    bool
	Start       engine.Tick
	Duration    engine.Tick
	Interval    engine.Tick
}

type EventManager struct {
	Events       []EventChance
	ActiveEvents map[EventID]*Event
	currentID    EventID
}

func NewEventManager() *EventManager {
	return &EventManager{
		ActiveEvents: make(map[EventID]*Event),
		Events: []EventChance{
			{0.01, NewPirateRaid},
		},
	}
}

func (em *EventManager) Add(event *Event, start engine.Tick) EventID {
	event.ID = em.currentID
	event.Start = start
	em.ActiveEvents[event.ID] = event
	em.currentID++

	return event.ID
}
