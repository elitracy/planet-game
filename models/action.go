package models

type Action interface {
	GetID() int
	GetTargetEntity() Entity
	GetDescription() string
	GetExecuteTick() int
	GetDuration() int
	GetStatus() EventStatus
	SetStatus(EventStatus)
	Execute()
}
