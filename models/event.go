package models

type EventStatus int

const (
	Pending = iota
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
