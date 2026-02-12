package actions

import (
	"fmt"

	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game/models"
)

type MoveEntityAction struct {
	*Action

	Destination models.Location
}

func NewMoveEntityAction(target models.Entity, dest models.Location, startTick engine.Tick, duration engine.Tick) *MoveEntityAction {

	action := &MoveEntityAction{
		Action: &Action{
			TargetEntity: target,
			Description:  fmt.Sprintf("Send %v to %v", target.GetName(), dest),
			StartTick:    startTick,
			Duration:     duration,
			Status:       engine.EventPending,
		},
		Destination: dest,
	}

	return action
}

func (a *MoveEntityAction) Execute() {
	if shipEntity, ok := a.TargetEntity.(*models.Ship); ok {
		shipEntity.Location = a.Destination
		engine.Info("%v: Moved Ship", shipEntity.GetName())
	}
}
