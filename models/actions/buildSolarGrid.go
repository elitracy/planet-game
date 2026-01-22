package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	 "github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
	 "github.com/elitracy/planets/core/state"
	 "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/constructions"
)

type BuildSolarGrid struct {
	ID           int
	TargetEntity models.Entity
	Description  string
	ExecuteTick  core.Tick
	Duration     core.Tick
	Status       consts.EventStatus
	Order        models.Order
}

func NewBuildSolarGridAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick, order models.Order) *BuildSolarGrid {

	action := &BuildSolarGrid{
		ID:           state.State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  fmt.Sprintf("Building a solar grid on %v", targetEntity.GetName()),
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       consts.EventPending,
		Order:        order,
	}

	return action
}

func (a BuildSolarGrid) GetID() int                    { return a.ID }
func (a BuildSolarGrid) GetTargetEntity() models.Entity       { return a.TargetEntity }
func (a BuildSolarGrid) GetDescription() string        { return a.Description }
func (a BuildSolarGrid) GetExecuteTick() core.Tick     { return a.ExecuteTick }
func (a BuildSolarGrid) GetDuration() core.Tick        { return a.Duration }
func (a BuildSolarGrid) GetStatus() consts.EventStatus        { return a.Status }
func (a *BuildSolarGrid) SetStatus(status consts.EventStatus) { a.Status = status }
func (a BuildSolarGrid) GetOrder() models.Order               { return a.Order }

func (a *BuildSolarGrid) Execute() {

	if planetEntity, ok := a.TargetEntity.(*models.Planet); ok {
		SolarGrid := constructions.CreateSolarGrid(1)
		planetEntity.Constructions.SolarGrids = append(planetEntity.Constructions.SolarGrids, SolarGrid)

		logging.Info("%v: Added SolarGrid", planetEntity.GetName())
	}
}
