package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	. "github.com/elitracy/planets/core/consts"
	. "github.com/elitracy/planets/core/interfaces"
	"github.com/elitracy/planets/core/logging"
	. "github.com/elitracy/planets/core/state"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/constructions"
)

type BuildMine struct {
	ID           int
	TargetEntity Entity
	Description  string
	ExecuteTick  core.Tick
	Duration     core.Tick
	Status       EventStatus
	Order        Order
}

func NewBuildMineAction(targetEntity Entity, executeTick core.Tick, duration core.Tick, order Order) *BuildMine {

	action := &BuildMine{
		ID:           State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  fmt.Sprintf("Building a mine on %v", targetEntity.GetName()),
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       Pending,
		Order:        order,
	}

	return action
}

func (a BuildMine) GetID() int                    { return a.ID }
func (a BuildMine) GetTargetEntity() Entity       { return a.TargetEntity }
func (a BuildMine) GetDescription() string        { return a.Description }
func (a BuildMine) GetExecuteTick() core.Tick     { return a.ExecuteTick }
func (a BuildMine) GetDuration() core.Tick        { return a.Duration }
func (a BuildMine) GetStatus() EventStatus        { return a.Status }
func (a *BuildMine) SetStatus(status EventStatus) { a.Status = status }
func (a BuildMine) GetOrder() Order               { return a.Order }
func (a *BuildMine) Execute() {
	if planetEntity, ok := a.TargetEntity.(*Planet); ok {
		Mine := constructions.CreateMine(1)
		planetEntity.Constructions.Mines = append(planetEntity.Constructions.Mines, Mine)

		logging.Info("%v: Added Mine", planetEntity.GetName())
	}
}
