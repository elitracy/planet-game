package constructions

const INITIAL_FOOD = 500

var farmConstructionTiers = map[int]int{
	1: 5,
	2: 10,
	3: 20,
}

type Farm struct {
	Quantity          int
	ProductionRate    int
	ConstructionTiers map[int]int
}

func CreateFarm(tier int) Farm {
	return Farm{
		Quantity:          INITIAL_FOOD,
		ProductionRate:    farmConstructionTiers[tier],
		ConstructionTiers: farmConstructionTiers,
	}
}

func (f *Farm) GetQuantity() int {
	return f.Quantity
}

func (f *Farm) GetProductionRate() int {
	return f.ProductionRate
}

func (f *Farm) GetTierRate(t int) int {
	return f.ConstructionTiers[t]
}
