package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/models"
)

type ScoutEntityAction struct {
	*Action
}

func NewScoutEntityAction(entity models.Entity, startTick engine.Tick, duration engine.Tick) *ScoutEntityAction {

	action := &ScoutEntityAction{
		Action: &Action{
			TargetEntity: entity,
			Description:  fmt.Sprintf("Survey %v", entity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       task.Pending,
		},
	}

	action.Action.Execute = func() {
		if system, ok := action.TargetEntity.(*models.StarSystem); ok {
			system.Scouted = true
		}

		if planet, ok := action.TargetEntity.(*models.Planet); ok {
			planet.Scouted = true
		}

	}

	return action
}
