package actions

import (
	"fmt"

	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/constructions"
	. "github.com/elitracy/planets/state"
)

type BuildFarm struct {
	ID           int
	TargetEntity Entity
	Description  string
	ExecuteTick  int
	Duration     int
	Status       EventStatus
	Order        Order
}

func NewBuildFarmAction(targetEntity Entity, executeTick int, duration int, order Order) *BuildFarm {

	action := &BuildFarm{
		ID:           State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  fmt.Sprintf("Building a farm on %v", targetEntity.GetName()),
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       Pending,
		Order:        order,
	}

	return action
}

func (a BuildFarm) GetID() int                    { return a.ID }
func (a BuildFarm) GetTargetEntity() Entity       { return a.TargetEntity }
func (a BuildFarm) GetDescription() string        { return a.Description }
func (a BuildFarm) GetExecuteTick() int           { return a.ExecuteTick }
func (a BuildFarm) GetDuration() int              { return a.Duration }
func (a BuildFarm) GetStatus() EventStatus        { return a.Status }
func (a *BuildFarm) SetStatus(status EventStatus) { a.Status = status }
func (a BuildFarm) GetOrder() Order               { return a.Order }

func (a *BuildFarm) Execute() {

	if planetEntity, ok := a.TargetEntity.(*Planet); ok {
		farm := constructions.CreateFarm(1)
		planetEntity.Constructions.Farms = append(planetEntity.Constructions.Farms, farm)

		logging.Info("%v: Added Farm", planetEntity.GetName())
	}
}
