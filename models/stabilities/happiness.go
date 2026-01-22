package stabilities

type Happiness struct {
	Quantity   float32
	GrowthRate float32
}

func (s *Happiness) GetQuantity() float32 {
	return s.Quantity
}

func (s *Happiness) GetGrowthRate() float32 {
	return s.GrowthRate
}

func (s *Happiness) Tick() {
	if s.GrowthRate > 0 {
		s.Quantity = min(1, s.Quantity+s.GrowthRate)
	}

	if s.GrowthRate < 0 {
		s.Quantity = max(-1, s.Quantity+s.GrowthRate)
	}
}
