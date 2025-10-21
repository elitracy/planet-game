package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	. "github.com/elitracy/planets/core/state"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/actions"
)

type CreateColony struct {
	ID          int
	Name        string
	Actions     []Action
	ExecuteTick core.Tick
	Status      consts.EventStatus
	*Planet
}

func NewCreateColonyOrder(planet *Planet, execTick core.Tick) *CreateColony {
	order := &CreateColony{
		ID:          State.OrderScheduler.GetNextID(),
		Name:        "Create Colony",
		ExecuteTick: execTick,
		Status:      consts.Pending,
		Planet:      planet,
	}

	createFarmAction := actions.NewBuildFarmAction(
		planet,
		order.ExecuteTick,
		40,
		order,
	)

	createMineAction := actions.NewBuildMineAction(
		planet,
		order.ExecuteTick,
		40,
		order,
	)

	createSolarGridAction := actions.NewBuildSolarGridAction(
		planet,
		order.ExecuteTick,
		80,
		order,
	)

	order.Actions = append(order.Actions, createFarmAction, createMineAction, createSolarGridAction)

	return order
}

func (o CreateColony) GetID() int                    { return o.ID }
func (o CreateColony) GetName() string               { return o.Name }
func (o CreateColony) GetActions() []Action          { return o.Actions }
func (o CreateColony) GetExecuteTick() core.Tick     { return o.ExecuteTick }
func (o CreateColony) GetStatus() consts.EventStatus { return o.Status }

func (o CreateColony) GetDuration() core.Tick {
	var duration core.Tick
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}

func (o *CreateColony) SetStatus(status consts.EventStatus) { o.Status = status }
