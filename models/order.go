package models

type Status int

const (
	Pending = iota
	Executing
	Complete
	Failed
)

type Order struct {
	ID            int
	TargetEntity  Entity
	Actions       []*Action // action queue, how to handle concurrent actions?
	StartTime     int
	StartPosition Position
	Status
	Velocity
}

func (o Order) GetETA() int {
	distance := int(EuclidianDistance(o.StartPosition, o.TargetEntity.GetPosition()))

	return distance / o.Velocity.Vector()
}

func (o Order) CurrentPosition(state *GameState) Position {
	// v = d / t
	// v = (p0-p1)/t
	// (p0 - vt) = p1
	timeElapsed := state.CurrentTick - o.StartTime
	vtx := o.Velocity.X * timeElapsed
	vty := o.Velocity.Y * timeElapsed
	vtz := o.Velocity.Z * timeElapsed

	currentPosX := o.StartPosition.X - vtx
	currentPosY := o.StartPosition.Y - vty
	currentPosZ := o.StartPosition.Z - vtz

	currentPosition := Position{currentPosX, currentPosY, currentPosZ}
	return currentPosition
}

// nil means no action executed
func (o *Order) ExecuteNextAction() *Action {
	for _, a := range o.Actions {
		if a.Status == Pending {
			a.Execute()
			return a
		}
	}
	return nil
}
