package resources

type Resource interface {
	GetName() string
	GetQuantity() int
	GetConsumptionRate() int // amount consumed per person
}
