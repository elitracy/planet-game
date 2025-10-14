package models

type ShipType int

const (
	Scout ShipType = iota
	Fighter
)

type Ship struct {
	ID         int
	Name       string
	OrderQueue []*Order
	Position
	ShipType
}

func (s Ship) GetID() int            { return s.ID }
func (s Ship) GetName() string       { return s.Name }
func (s Ship) GetPosition() Position { return s.Position }
func (s Ship) GetOrders() []*Order   { return s.OrderQueue }

type ShipManager struct {
	Ships     map[int]*Ship
	currentID int
}

func (m ShipManager) GetShip(id int) *Ship { return m.Ships[id] }

func (m *ShipManager) CreateShip(name string, position Position, shipType ShipType) *Ship {
	ship := &Ship{
		ID:       m.GetNextID(),
		Name:     name,
		Position: position,
		ShipType: shipType,
	}

	m.Ships[ship.GetID()] = ship

	return ship
}

func (m *ShipManager) GetNextID() int {
	m.currentID++
	return m.currentID
}
