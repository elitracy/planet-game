package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/constructions"
	"github.com/elitracy/planets/models/events"
)

type BuildSolarGridAction struct {
	*Action
}

func NewBuildSolarGridAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick) *BuildSolarGridAction {

	action := &BuildSolarGridAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Building a solar grid on %v", targetEntity.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       events.EventPending,
		},
	}

	return action
}

func (a *BuildSolarGridAction) Execute() {

	if planetEntity, ok := a.Action.TargetEntity.(*models.Planet); ok {
		SolarGrid := constructions.CreateSolarGrid(1)
		planetEntity.Constructions.SolarGrids = append(planetEntity.Constructions.SolarGrids, SolarGrid)

		logging.Info("%v: Added SolarGrid", planetEntity.GetName())
	}
}
