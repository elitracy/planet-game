package stabilities

type Corruption struct {
	Quantity   float32
	GrowthRate float32
}

func (s *Corruption) GetQuantity() float32 {
	return s.Quantity
}

func (s *Corruption) GetGrowthRate() float32 {
	return s.GrowthRate
}

func (s *Corruption) Tick() {
	if s.GrowthRate > 0 {
		s.Quantity = min(1, s.Quantity+s.GrowthRate)
	}

	if s.GrowthRate < 0 {
		s.Quantity = max(-1, s.Quantity+s.GrowthRate)
	}
}
