package resources

type Energy struct {
	Quantity        int
	ConsumptionRate int
}

func (e Energy) GetQuantity() int {
	return e.Quantity
}

func (e Energy) GetConsumptionRate() int {
	return e.ConsumptionRate
}
