package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
)

type ScoutSystemAction struct{
	*Action
}


func NewScoutSystemAction(system *models.StarSystem, executeTick core.Tick, duration core.Tick) *TimeoutAction {

	action := &TimeoutAction{
		Action: &Action{
			TargetEntity: system,
			Description:  fmt.Sprintf("Scoutting %v", system.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       consts.EventPending,
		},
	}

	return action
}

func (a *ScoutSystemAction)Execute(){
	if system, ok := a.TargetEntity.(*models.StarSystem); ok {
		system.Scouted = true
	}
}
