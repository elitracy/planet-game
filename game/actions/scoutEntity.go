package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
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
			Status:       engine.EventPending,
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
