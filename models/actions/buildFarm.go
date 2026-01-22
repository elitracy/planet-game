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

type BuildFarm struct {
	ID           int
	TargetEntity models.Entity
	Description  string
	ExecuteTick  core.Tick
	Duration     core.Tick
	Status       consts.EventStatus
	Order        models.Order
}

func NewBuildFarmAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick, order models.Order) *BuildFarm {

	action := &BuildFarm{
		ID:           state.State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  fmt.Sprintf("Building a farm on %v", targetEntity.GetName()),
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       consts.EventPending,
		Order:        order,
	}

	return action
}

func (a BuildFarm) GetID() int                           { return a.ID }
func (a BuildFarm) GetTargetEntity() models.Entity       { return a.TargetEntity }
func (a BuildFarm) GetDescription() string               { return a.Description }
func (a BuildFarm) GetExecuteTick() core.Tick            { return a.ExecuteTick }
func (a BuildFarm) GetDuration() core.Tick               { return a.Duration }
func (a BuildFarm) GetStatus() consts.EventStatus        { return a.Status }
func (a *BuildFarm) SetStatus(status consts.EventStatus) { a.Status = status }
func (a BuildFarm) GetOrder() models.Order               { return a.Order }

func (a *BuildFarm) Execute() {

	if planetEntity, ok := a.TargetEntity.(*models.Planet); ok {
		farm := constructions.CreateFarm(1)
		planetEntity.Constructions.Farms = append(planetEntity.Constructions.Farms, farm)

		logging.Info("%v: Added Farm", planetEntity.GetName())
	}
}
