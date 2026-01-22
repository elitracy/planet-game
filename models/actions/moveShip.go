package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	 "github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
	 "github.com/elitracy/planets/core/state"
	 "github.com/elitracy/planets/models"
)

type MoveShip struct {
	ID           int
	TargetEntity models.Entity
	Description  string
	ExecuteTick  core.Tick
	Duration     core.Tick
	Status       consts.EventStatus

	Destination core.Position
}

func NewMoveShipAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick, destination core.Position) *MoveShip {

	action := &MoveShip{
		ID:           state.State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  fmt.Sprintf("Sending a ship [%v] to %v", targetEntity.GetName(), destination),
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       consts.EventPending,
		Destination:  destination,
	}

	return action
}

func (a MoveShip) GetID() int                    { return a.ID }
func (a MoveShip) GetTargetEntity() models.Entity       { return a.TargetEntity }
func (a MoveShip) GetDescription() string        { return a.Description }
func (a MoveShip) GetExecuteTick() core.Tick     { return a.ExecuteTick }
func (a MoveShip) GetDuration() core.Tick        { return a.Duration }
func (a MoveShip) GetStatus() consts.EventStatus        { return a.Status }
func (a *MoveShip) SetStatus(status consts.EventStatus) { a.Status = status }
func (a MoveShip) GetDestination() core.Position      { return a.Destination }

func (a *MoveShip) Execute() {
	if shipEntity, ok := a.TargetEntity.(*models.Ship); ok {
		shipEntity.Position = a.Destination
		logging.Info("%v: Moved Ship", shipEntity.GetName())
	}
}
