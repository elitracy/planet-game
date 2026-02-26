package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/models"
)

type ColonizeAction struct {
	*Action
	OnColonize func()
}

func NewColonizeAction(targetEntity models.Entity, startTick engine.Tick, duration engine.Tick, onExecute func()) *ColonizeAction {

	action := &ColonizeAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Colonize %v", targetEntity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       task.Pending,
			Execute:      onExecute,
		},
	}

	return action
}
