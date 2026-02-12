package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/models"
)

type TimeoutAction struct {
	*Action
}

func NewTimeoutAction(targetEntity models.Entity, startTick engine.Tick, duration engine.Tick, destination models.Entity) *TimeoutAction {

	action := &TimeoutAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Sending info about %v", targetEntity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       engine.EventPending,
		},
	}

	return action
}

func (a *TimeoutAction) Execute() { return }
