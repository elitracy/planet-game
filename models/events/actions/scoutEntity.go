package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
)

type ScoutEntityAction struct {
	*Action
}

func NewScoutEntityAction(entity models.Entity, executeTick core.Tick, duration core.Tick) *ScoutEntityAction{

	action := &ScoutEntityAction{
		Action: &Action{
			TargetEntity: entity,
			Description:  fmt.Sprintf("Scoutting %v", entity.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       events.EventPending,
		},
	}

	return action
}

func (a *ScoutEntityAction) Execute() {
	if system, ok := a.TargetEntity.(*models.StarSystem); ok {
		system.Scouted = true
	}

	if planet, ok := a.TargetEntity.(*models.Planet); ok {
		planet.Scouted = true
	}
}
