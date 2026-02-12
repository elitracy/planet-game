package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/models/constructions"
	"github.com/elitracy/planets/game/models"
)

type BuildFarmAction struct {
	*Action
}

func NewBuildFarmAction(targetEntity models.Entity, startTick engine.Tick, duration engine.Tick) *BuildFarmAction {

	action := &BuildFarmAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Building a farm on %v", targetEntity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       engine.EventPending,
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
