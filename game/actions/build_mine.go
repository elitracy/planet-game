package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/models"
	"github.com/elitracy/planets/game/models/constructions"
)

type BuildMineAction struct {
	*Action
}

func NewBuildMineAction(targetEntity models.Entity, startTick engine.Tick, duration engine.Tick) *BuildMineAction {

	action := &BuildMineAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Building a mine on %v", targetEntity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       task.Pending,
		},
	}
	action.Action.Execute = func() {
		if planetEntity, ok := action.Action.TargetEntity.(*models.Planet); ok {
			Mine := constructions.CreateMine(1)
			planetEntity.Constructions.Mines = append(planetEntity.Constructions.Mines, Mine)

			engine.Info("%v: Added Mine", planetEntity.GetName())
		}
	}

	return action
}
