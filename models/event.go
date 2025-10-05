package models

//go:generate stringer -type=EventStatus
type EventStatus int

const (
	Pending EventStatus = iota
	Executing
	Complete
	Failed
)

type Event interface {
	GetID() int
	GetStart() int
	GetDuration() int
	GetStatus() EventStatus
}
