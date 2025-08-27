package resources

type Mineral struct {
	Quantity        int
	ConsumptionRate int
}

func (f *Mineral) GetQuantity() int {
	return f.Quantity
}

func (f *Mineral) GetConsumptionRate() int {
	return f.ConsumptionRate
}
