package stabilities

type Happiness struct {
	Name     string
	Quantity float32
}

func (h *Happiness) GetName() string {
	return h.Name
}

func (h *Happiness) GetQuantity() float32 {
	return h.Quantity
}
