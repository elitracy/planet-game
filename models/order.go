package models

type Order interface {
	GetID() int
	GetName() string
	GetExecuteTick() int
	GetDuration() int
	GetStatus() EventStatus
	SetStatus(EventStatus)

	GetActions() []Action
}
