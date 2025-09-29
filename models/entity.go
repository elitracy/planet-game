package models

type Entity interface {
	GetID() int
	GetPosition() Position
	GetOrders() []Order
}
