package stabilities

type Corruption struct {
	Name     string
	Quantity float32
}

func (c *Corruption) GetName() string {
	return c.Name
}

func (c *Corruption) GetQuantity() float32 {
	return c.Quantity
}
