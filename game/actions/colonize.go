package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/models"
)

type ColonizeAction struct {
	*Action
}

func NewColonizeAction(targetEntity models.Entity, startTick engine.Tick, duration engine.Tick) *ColonizeAction {

	action := &ColonizeAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Colonize %v", targetEntity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       task.Pending,
		},
	}

	action.Action.Execute = func() {
		if system, ok := action.Action.TargetEntity.(*models.StarSystem); ok {
			system.Colonized = true
			return
		}

		if planet, ok := action.Action.TargetEntity.(*models.Planet); ok {
			planet.Colonized = true
			return
		}
	}
	return action
}
