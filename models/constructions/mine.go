package constructions

type Mine struct {
	Name     string
	Quantity int
}

func (m *Mine) GetName() string {
	return m.Name
}

func (m *Mine) GetQuantity() int {
	return m.Quantity
}
