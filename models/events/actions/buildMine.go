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

func NewBuildMineAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick) *BuildMineAction {

	action := &BuildMineAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Building a mine on %v", targetEntity.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration, // TODO: eventually Tick * 40 for clarity
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
