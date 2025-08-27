package constructions

type Mine struct {
	Quantity       int
	ProductionRate int
}

func (m *Mine) GetQuantity() int {
	return m.Quantity
}

func (m *Mine) GetProductionRate() int {
	return m.ProductionRate
}
