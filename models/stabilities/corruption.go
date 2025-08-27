package stabilities

type Corruption struct {
	Quantity   float32
	GrowthRate float32
}

func (c *Corruption) GetQuantity() float32 {
	return c.Quantity
}

func (c *Corruption) GetGrowthRate() float32 {
	return c.GrowthRate
}
