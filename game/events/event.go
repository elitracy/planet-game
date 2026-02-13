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

type Event struct {
	ID          EventID
	Name        string
	Description string
	Severity    EventSeverity
	Target      models.Entity
	Deadline    engine.Tick
	Effect      func()
	Resolve     func()
	Expire      func()
	Resolved    bool
}
