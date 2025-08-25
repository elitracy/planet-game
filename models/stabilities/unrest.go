package stabilities

type Unrest struct {
	Name     string
	Quantity float32
}

func (u *Unrest) GetName() string {
	return u.Name
}

func (u *Unrest) GetQuantity() float32 {
	return u.Quantity
}
