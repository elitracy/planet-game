package resources

type Food struct {
	Quantity        int
	ConsumptionRate int
}

func (f *Food) GetQuantity() int {
	return f.Quantity
}

func (f *Food) GetConsumptionRate() int {
	return f.ConsumptionRate
}
