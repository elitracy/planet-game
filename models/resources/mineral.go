package resources

type Mineral struct {
	Name            string
	Quantity        int
	ConsumptionRate int
}

func (f *Mineral) GetName() string {
	return f.Name
}

func (f *Mineral) GetQuantity() int {
	return f.Quantity
}

func (f *Mineral) GetConsumptionRate() int {
	return f.ConsumptionRate
}
