package constructions

type Construction interface {
	GetQuantity() int
	GetProductionRate() int
	GetTierRate(int) int
}
