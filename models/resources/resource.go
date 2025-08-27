package resources

type Resource interface {
	GetQuantity() int
	GetConsumptionRate() int // amount consumed per person
}
