package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/models"
	"github.com/elitracy/planets/game/models/constructions"
)

type BuildSolarGridAction struct {
	*Action
}

func NewBuildSolarGridAction(targetEntity models.Entity, startTick engine.Tick, duration engine.Tick) *BuildSolarGridAction {

	action := &BuildSolarGridAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Building a solar grid on %v", targetEntity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       task.Pending,
		},
	}

	action.Action.Execute = func() {

		if planetEntity, ok := action.Action.TargetEntity.(*models.Planet); ok {
			SolarGrid := constructions.CreateSolarGrid(1)
			planetEntity.Constructions.SolarGrids = append(planetEntity.Constructions.SolarGrids, SolarGrid)

			engine.Info("%v: Added SolarGrid", planetEntity.GetName())
		}
	}

	return action
}
