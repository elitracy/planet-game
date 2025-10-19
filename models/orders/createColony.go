package orders

import (
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/actions"
	. "github.com/elitracy/planets/state"
)

type CreateColony struct {
	ID          int
	Name        string
	Actions     []Action
	ExecuteTick int
	Status      EventStatus
	*Planet
}

func NewCreateColonyOrder(planet *Planet, execTick int) *CreateColony {
	order := &CreateColony{
		ID:          State.OrderScheduler.GetNextID(),
		Name:        "Create Colony",
		ExecuteTick: execTick,
		Status:      Pending,
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

func (o CreateColony) GetID() int             { return o.ID }
func (o CreateColony) GetName() string        { return o.Name }
func (o CreateColony) GetActions() []Action   { return o.Actions }
func (o CreateColony) GetExecuteTick() int    { return o.ExecuteTick }
func (o CreateColony) GetStatus() EventStatus { return o.Status }

func (o CreateColony) GetDuration() int {
	duration := 0
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}

func (o *CreateColony) SetStatus(status EventStatus) { o.Status = status }
