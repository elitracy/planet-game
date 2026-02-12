package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
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
			Status:       engine.EventPending,
		},
	}

	return action
}

func (a *ColonizeAction) Execute() {

	if system, ok := a.Action.TargetEntity.(*models.StarSystem); ok {
		system.Colonized = true
		return
	}

	if planet, ok := a.Action.TargetEntity.(*models.Planet); ok {
		planet.Colonized = true
		return
	}
}
