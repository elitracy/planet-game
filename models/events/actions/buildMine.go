package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/events"
)

type BuildMineAction struct {
	*Action
}

func NewBuildMineAction(targetEntity models.Entity, startTick core.Tick, duration core.Tick) *BuildMineAction {

	action := &BuildMineAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Building a mine on %v", targetEntity.GetName()),
			StartTick:    startTick,
			Duration:     duration,
			Status:       events.EventPending,
		},
	}

	return action
}

func (a *BuildMineAction) Execute() {
	if planetEntity, ok := a.Action.TargetEntity.(*models.Planet); ok {
		Mine := constructions.CreateMine(1)
		planetEntity.Constructions.Mines = append(planetEntity.Constructions.Mines, Mine)

		logging.Info("%v: Added Mine", planetEntity.GetName())
	}
}
