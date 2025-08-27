package constructions

type Farm struct {
	Quantity       int
	ProductionRate int
}

func (f *Farm) GetQuantity() int {
	return f.Quantity
}

func (f *Farm) GetProductionRate() int {
	return f.ProductionRate
}
