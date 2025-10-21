package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/consts"
	. "github.com/elitracy/planets/core/interfaces"
	"github.com/elitracy/planets/core/logging"
	. "github.com/elitracy/planets/core/state"
	. "github.com/elitracy/planets/models"
)

type MoveShip struct {
	ID           int
	TargetEntity Entity
	Description  string
	ExecuteTick  core.Tick
	Duration     core.Tick
	Status       EventStatus

	Destination Position
}

func NewMoveShipAction(targetEntity Entity, executeTick core.Tick, duration core.Tick, destination Position) *MoveShip {

	action := &MoveShip{
		ID:           State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  fmt.Sprintf("Sending a ship [%v] to %v", targetEntity.GetName(), destination),
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       Pending,
		Destination:  destination,
	}

	return action
}

func (a MoveShip) GetID() int                    { return a.ID }
func (a MoveShip) GetTargetEntity() Entity       { return a.TargetEntity }
func (a MoveShip) GetDescription() string        { return a.Description }
func (a MoveShip) GetExecuteTick() core.Tick     { return a.ExecuteTick }
func (a MoveShip) GetDuration() core.Tick        { return a.Duration }
func (a MoveShip) GetStatus() EventStatus        { return a.Status }
func (a *MoveShip) SetStatus(status EventStatus) { a.Status = status }
func (a MoveShip) GetDestination() Position      { return a.Destination }

func (a *MoveShip) Execute() {
	if shipEntity, ok := a.TargetEntity.(*Ship); ok {
		shipEntity.Position = a.Destination
		logging.Info("%v: Moved Ship", shipEntity.GetName())
	}
}
