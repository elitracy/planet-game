package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/config"
	"github.com/elitracy/planets/game/models"
)

type TimeoutAction struct {
	*Action
}

func NewTimeoutAction(targetEntity models.Entity,  destination models.Entity, startTick engine.Tick, duration engine.Tick) *TimeoutAction {

	action := &TimeoutAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Waiting %v", config.FormatGameTime(duration)),
			StartTick:    startTick,
			Duration:     duration,
			Status:       engine.EventPending,
		},
	}

	return action
}
