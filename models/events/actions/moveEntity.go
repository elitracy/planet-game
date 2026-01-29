package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/events"
)

type MoveEntityAction struct {
	*Action

	Destination models.Location
}

func NewMoveEntityAction(target models.Entity, dest models.Location, executeTick core.Tick, duration core.Tick) *MoveEntityAction {

	action := &MoveEntityAction{
		Action: &Action{
			TargetEntity: target,
			Description:  fmt.Sprintf("Sending %v to %v", target.GetName(), dest),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       events.EventPending,
		},
		Destination: dest,
	}

	return action
}

func (a *MoveEntityAction) Execute() {
	if shipEntity, ok := a.TargetEntity.(*models.Ship); ok {
		shipEntity.Location = a.Destination
		logging.Info("%v: Moved Ship", shipEntity.GetName())
	}
}
