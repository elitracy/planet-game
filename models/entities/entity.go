package models

type Entity interface {
	GetID()
	GetPosition()
	GetOrders()
}
