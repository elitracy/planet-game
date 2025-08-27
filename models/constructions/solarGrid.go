package constructions

type SolarGrid struct {
	Quantity       int
	ProductionRate int
}

func (sg *SolarGrid) GetQuantity() int {
	return sg.Quantity
}

func (sg *SolarGrid) GetProductionRate() int {
	return sg.ProductionRate
}
