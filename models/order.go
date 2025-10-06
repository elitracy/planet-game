package models

//go:generate stringer -type=OrderType
type OrderType int

const (
	CreateColonyOrder OrderType = iota
	SendFoodOrder
	SendMineralOrder
	SendEnergyOrder
)

type Order struct {
	ID            int
	TargetEntity  Entity
	Type          OrderType
	Actions       []*Action
	ExecuteTime   int
	StartPosition Position
	Status        EventStatus
	// Velocity
}

func CreateNewOrder(entity Entity, orderType OrderType, start int, pos Position) *Order {
	order := &Order{
		ID:            GameStateGlobal.OrderScheduler.GetNextID(),
		TargetEntity:  entity,
		Type:          orderType,
		ExecuteTime:   start,
		StartPosition: pos,
		Status:        Pending,
	}

	if order.Type == CreateColonyOrder {

		createFarmAction := &Action{
			ID:           GameStateGlobal.ActionScheduler.GetNextID(),
			TargetEntity: entity,
			Description:  "Builds a farm on the target Planet",
			Type:         BuildFarm,
			ExecuteTime:  order.ExecuteTime + 5,
			Status:       Pending,
			Order:        order,
		}

		createMineAction := &Action{
			ID:           GameStateGlobal.ActionScheduler.GetNextID(),
			TargetEntity: entity,
			Description:  "Builds a mine on the target Planet",
			Type:         BuildMine,
			ExecuteTime:  order.ExecuteTime + 5,
			Status:       Pending,
			Order:        order,
		}

		createSolarGridAction := &Action{
			ID:           GameStateGlobal.ActionScheduler.GetNextID(),
			TargetEntity: entity,
			Description:  "Builds a solar grid on the target Planet",
			Type:         BuildSolarGrid,
			ExecuteTime:  order.ExecuteTime + 10,
			Status:       Pending,
			Order:        order,
		}

		order.Actions = append(order.Actions, createFarmAction, createMineAction, createSolarGridAction)
	}

	return order
}

func (o Order) GetID() int {
	return o.ID
}

func (o Order) GetStart() int {
	return o.ExecuteTime
}

func (o Order) GetDuration() int {
	duration := 0
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}

	return duration
}

func (o Order) GetStatus() EventStatus {
	return o.Status
}

// func (o Order) GetETA() int {
// 	distance := int(EuclidianDistance(o.StartPosition, o.TargetEntity.GetPosition()))
//
// 	return distance / o.Velocity.Vector()
// }

// func (o Order) CurrentPosition(state *GameState) Position {
// 	// v = d / t
// 	// v = (p0-p1)/t
// 	// (p0 - vt) = p1
// 	timeElapsed := state.CurrentTick - o.StartTime
// 	vtx := o.Velocity.X * timeElapsed
// 	vty := o.Velocity.Y * timeElapsed
// 	vtz := o.Velocity.Z * timeElapsed
//
// 	currentPosX := o.StartPosition.X - vtx
// 	currentPosY := o.StartPosition.Y - vty
// 	currentPosZ := o.StartPosition.Z - vtz
//
// 	currentPosition := Position{currentPosX, currentPosY, currentPosZ}
// 	return currentPosition
// }

// nil means no action executed
