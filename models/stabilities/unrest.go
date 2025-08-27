package stabilities

type Unrest struct {
	Quantity   float32
	GrowthRate float32
}

func (u *Unrest) GetQuantity() float32 {
	return u.Quantity
}

func (u *Unrest) GetGrowthRate() float32 {
	return u.GrowthRate
}
