package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/models"
	"github.com/elitracy/planets/game/models/constructions"
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
			Status:       task.Pending,
		},
	}

	action.Action.Execute = func() {
		if planetEntity, ok := action.TargetEntity.(*models.Planet); ok {
			farm := constructions.CreateFarm(1)
			planetEntity.Constructions.Farms = append(planetEntity.Constructions.Farms, farm)
		}

	}

	return action
}
