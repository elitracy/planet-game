package resources

type Food struct {
	Name            string
	Quantity        int
	ConsumptionRate int
}

func (f *Food) GetName() string {
	return f.Name
}

func (f *Food) GetQuantity() int {
	return f.Quantity
}

func (f *Food) GetConsumptionRate() int {
	return f.ConsumptionRate
}
