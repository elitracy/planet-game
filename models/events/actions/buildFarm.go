package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/events"
)

type BuildFarmAction struct {
	*Action
}

func NewBuildFarmAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick) *BuildFarmAction {

	action := &BuildFarmAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Building a farm on %v", targetEntity.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       events.EventPending,
		},
	}

	return action
}

func (a *BuildFarmAction) Execute() {

	if planetEntity, ok := a.Action.TargetEntity.(*models.Planet); ok {
		farm := constructions.CreateFarm(1)
		planetEntity.Constructions.Farms = append(planetEntity.Constructions.Farms, farm)
	}
}
