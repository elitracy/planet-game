package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
)

type ColonizeAction struct {
	*Action
}

func NewColonizeAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick) *ColonizeAction {

	action := &ColonizeAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Colonize %v", targetEntity.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       events.EventPending,
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
