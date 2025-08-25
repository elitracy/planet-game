package stabilities

type Stability interface {
	GetName() string
	GetQuantity() float32
}
