package orders

import . "github.com/elitracy/planets/models"

type CreateColonyOrder struct {
	ID          int
	Name        string
	Actions     []*Action
	ExecuteTick int
	Status      EventStatus
	*Planet
}

func NewCreateColonyOrder(planet *Planet, execTick int) *CreateColonyOrder {
	order := &CreateColonyOrder{
		ID:          GameStateGlobal.OrderScheduler.GetNextID(),
		Name:        "Create Colony",
		ExecuteTick: execTick,
		Status:      Pending,
		Planet:      planet,
	}

	createFarmAction := &Action{
		ID:           GameStateGlobal.ActionScheduler.GetNextID(),
		TargetEntity: planet,
		Description:  "Builds a farm on the target Planet",
		Type:         BuildFarm,
		ExecuteTick:  order.ExecuteTick,
		Duration:     40, // TODO: eventually Tick * 40 for clarity
		Status:       Pending,
		Order:        order,
	}

	createMineAction := &Action{
		ID:           GameStateGlobal.ActionScheduler.GetNextID(),
		TargetEntity: planet,
		Description:  "Builds a mine on the target Planet",
		Type:         BuildMine,
		ExecuteTick:  order.ExecuteTick,
		Duration:     40,
		Status:       Pending,
		Order:        order,
	}

	createSolarGridAction := &Action{
		ID:           GameStateGlobal.ActionScheduler.GetNextID(),
		TargetEntity: planet,
		Description:  "Builds a solar grid on the target Planet",
		Type:         BuildSolarGrid,
		ExecuteTick:  order.ExecuteTick,
		Duration:     80,
		Status:       Pending,
		Order:        order,
	}

	order.Actions = append(order.Actions, createFarmAction, createMineAction, createSolarGridAction)

	return order
}

func (o CreateColonyOrder) GetID() int             { return o.ID }
func (o CreateColonyOrder) GetName() string        { return o.Name }
func (o CreateColonyOrder) GetActions() []*Action  { return o.Actions }
func (o CreateColonyOrder) GetExecuteTick() int    { return o.ExecuteTick }
func (o CreateColonyOrder) GetStatus() EventStatus { return o.Status }

func (o CreateColonyOrder) GetDuration() int {
	duration := 0
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}

func (o *CreateColonyOrder) SetStatus(status EventStatus) { o.Status = status }
