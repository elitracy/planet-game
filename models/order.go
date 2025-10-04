package models

type Order struct {
	ID            int
	TargetEntity  *Entity
	Actions       []*Action // action queue, how to handle concurrent actions?
	StartTime     int
	StartPosition Position
	Status        EventStatus
	// Velocity
}

func (o Order) GetID() int {
	return o.ID
}

func (o Order) GetStart() int {
	return o.StartTime
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
