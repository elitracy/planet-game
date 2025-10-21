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

type BuildSolarGrid struct {
	ID           int
	TargetEntity Entity
	Description  string
	ExecuteTick  core.Tick
	Duration     core.Tick
	Status       EventStatus
	Order        Order
}

func NewBuildSolarGridAction(targetEntity Entity, executeTick core.Tick, duration core.Tick, order Order) *BuildSolarGrid {

	action := &BuildSolarGrid{
		ID:           State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  fmt.Sprintf("Building a solar grid on %v", targetEntity.GetName()),
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       Pending,
		Order:        order,
	}

	return action
}

func (a BuildSolarGrid) GetID() int                    { return a.ID }
func (a BuildSolarGrid) GetTargetEntity() Entity       { return a.TargetEntity }
func (a BuildSolarGrid) GetDescription() string        { return a.Description }
func (a BuildSolarGrid) GetExecuteTick() core.Tick     { return a.ExecuteTick }
func (a BuildSolarGrid) GetDuration() core.Tick        { return a.Duration }
func (a BuildSolarGrid) GetStatus() EventStatus        { return a.Status }
func (a *BuildSolarGrid) SetStatus(status EventStatus) { a.Status = status }
func (a BuildSolarGrid) GetOrder() Order               { return a.Order }

func (a *BuildSolarGrid) Execute() {

	if planetEntity, ok := a.TargetEntity.(*Planet); ok {
		SolarGrid := constructions.CreateSolarGrid(1)
		planetEntity.Constructions.SolarGrids = append(planetEntity.Constructions.SolarGrids, SolarGrid)

		logging.Info("%v: Added SolarGrid", planetEntity.GetName())
	}
}
