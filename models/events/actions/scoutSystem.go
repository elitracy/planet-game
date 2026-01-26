package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
)

type ScoutSystemAction struct {
	*Action
}

func NewScoutSystemAction(system *models.StarSystem, executeTick core.Tick, duration core.Tick) *TimeoutAction {

	action := &TimeoutAction{
		Action: &Action{
			TargetEntity: system,
			Description:  fmt.Sprintf("Scoutting %v", system.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       events.EventPending,
		},
	}

	return action
}

func (a *ScoutSystemAction) Execute() {
	if system, ok := a.TargetEntity.(*models.StarSystem); ok {
		system.Scouted = true
	}
}
