package resources

type Mineral struct {
	Quantity        int
	ConsumptionRate int
}

func (m Mineral) GetQuantity() int {
	return m.Quantity
}

func (m Mineral) GetConsumptionRate() int {
	return m.ConsumptionRate
}
