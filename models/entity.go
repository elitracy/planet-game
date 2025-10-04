package models

type Entity interface {
	GetID() int
	GetName() string
	GetPosition() Position
	GetOrders() []Order
}
