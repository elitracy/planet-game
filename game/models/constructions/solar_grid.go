package constructions

const INITIAL_ENERGY = 500

var solarGridConstructionTiers = map[int]int{
	1: 5,
	2: 10,
	3: 20,
}

type SolarGrid struct {
	Quantity          int
	ProductionRate    int
	ConstructionTiers map[int]int
}

func CreateSolarGrid(tier int) SolarGrid {
	return SolarGrid{
		Quantity:          INITIAL_ENERGY,
		ProductionRate:    solarGridConstructionTiers[tier],
		ConstructionTiers: solarGridConstructionTiers,
	}
}

func (sg *SolarGrid) GetQuantity() int {
	return sg.Quantity
}

func (sg *SolarGrid) GetProductionRate() int {
	return sg.ProductionRate
}

func (sg *SolarGrid) GetTierRate(t int) int {
	return sg.ConstructionTiers[t]
}
