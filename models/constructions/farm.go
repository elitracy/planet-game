package constructions

type Farm struct {
	Name     string
	Quantity int
}

func (f *Farm) GetName() string {
	return f.Name
}

func (f *Farm) GetQuantity() int {
	return f.Quantity
}
