package events

import (
	"math/rand"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
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

const (
	minEventInterval = 5
	maxEventInterval = 10
)

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
	NextEvent    engine.Tick
}

func NewEventManager() *EventManager {
	return &EventManager{
		ActiveEvents: make(map[EventID]*Event),
		Events: []EventChance{
			{0.5, NewPirateRaid},
			{0.5, NewMineralDiscovery},
			{0.5, NewPlagueDisaster},
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

func (em *EventManager) RollNextEventInterval(currentTick engine.Tick) {

	interval := minEventInterval + engine.Tick(rand.Intn(maxEventInterval-minEventInterval))
	em.NextEvent = currentTick + (interval * config.TICKS_PER_PULSE)

	engine.Info("Next Event: %v", config.FormatGameTime(em.NextEvent))
}
