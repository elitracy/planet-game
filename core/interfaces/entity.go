package interfaces

import "github.com/elitracy/planets/core"

type Entity interface {
	GetID() int
	GetName() string
	GetPosition() core.Position
	GetOrders() []Event
}
