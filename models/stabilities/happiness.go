package stabilities

type Happiness struct {
	Quantity   float32
	GrowthRate float32
}

func (h *Happiness) GetQuantity() float32 {
	return h.Quantity
}

func (h *Happiness) GetGrowthRate() float32 {
	return h.GrowthRate
}
