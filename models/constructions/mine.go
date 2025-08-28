package constructions

const INITIAL_MINERAL = 500

var mineConstructionTiers = map[int]int{
	1: 5,
	2: 10,
	3: 20,
}

type Mine struct {
	Quantity          int
	ProductionRate    int
	ConstructionTiers map[int]int
}

func CreateMine(tier int) Mine {
	return Mine{
		Quantity:          INITIAL_MINERAL,
		ProductionRate:    mineConstructionTiers[tier],
		ConstructionTiers: mineConstructionTiers,
	}
}

func (m *Mine) GetQuantity() int {
	return m.Quantity
}

func (m *Mine) GetProductionRate() int {
	return m.ProductionRate
}

func (m *Mine) GetTierRate(t int) int {
	return m.ConstructionTiers[t]
}
