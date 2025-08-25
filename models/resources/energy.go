package resources

type Energy struct {
	Name            string
	Quantity        int
	ConsumptionRate int
}

func (e *Energy) GetName() string {
	return e.Name
}

func (e *Energy) GetQuantity() int {
	return e.Quantity
}

func (e *Energy) GetConsumptionRate() int {
	return e.ConsumptionRate
}
