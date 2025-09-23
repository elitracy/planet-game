package models

import models "github.com/elitracy/planets/models/entities"

type OrderStatus int

const (
	Pending = iota
	Executing
	Complete
	Failed
)

type Order interface {
	GetID() int
	GetTargetEntity() models.Entity
	GetAction() Action
	GetETA() int // tick
	GetStatus() OrderStatus
}
