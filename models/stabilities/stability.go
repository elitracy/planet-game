package stabilities

type Stability interface {
	GetQuantity() float32
	GetGrowthRate() float32
}
