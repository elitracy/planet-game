package constructions

type SolarGrid struct {
	Name     string
	Quantity int
}

func (sg *SolarGrid) GetName() string {
	return sg.Name
}

func (sg *SolarGrid) GetQuantity() int {
	return sg.Quantity
}
