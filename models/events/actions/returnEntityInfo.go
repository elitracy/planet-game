package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/models"
)

type TimeoutAction struct {
	*Action
}

func NewTimeoutAction(targetEntity models.Entity, executeTick core.Tick, duration core.Tick, destination models.Entity) *TimeoutAction {

	action := &TimeoutAction{
		Action: &Action{
			TargetEntity: targetEntity,
			Description:  fmt.Sprintf("Sending info about %v", targetEntity.GetName()),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       consts.EventPending,
		},
	}

	return action
}

func (a *TimeoutAction) Execute() { return }
