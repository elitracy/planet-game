package actions

import (
	"fmt"

	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models"
)

type MoveEntityAction struct {
	*Action

	Destination models.Destination
}

func NewMoveEntityAction(target models.Entity, dest models.Destination, executeTick core.Tick, duration core.Tick) *MoveEntityAction {

	action := &MoveEntityAction{
		Action: &Action{
			TargetEntity: target,
			Description:  fmt.Sprintf("Sending %v to %v", target.GetName(), dest),
			ExecuteTick:  executeTick,
			Duration:     duration,
			Status:       consts.EventPending,
		},
		Destination: dest,
	}

	return action
}

func (a *MoveEntityAction) Execute() {
	if shipEntity, ok := a.TargetEntity.(*models.Ship); ok {
		shipEntity.Position = a.Destination.Position
		logging.Info("%v: Moved Ship", shipEntity.GetName())
	}
}
