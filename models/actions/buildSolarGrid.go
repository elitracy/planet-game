package actions

import (
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/constructions"
	. "github.com/elitracy/planets/state"
)

type BuildSolarGrid struct {
	ID           int
	TargetEntity Entity
	Description  string
	ExecuteTick  int
	Duration     int
	Status       EventStatus
	Order        Order
}

func NewBuildSolarGridAction(targetEntity Entity, executeTick int, duration int, order Order) *BuildSolarGrid {

	action := &BuildSolarGrid{
		ID:           State.ActionScheduler.GetNextID(),
		TargetEntity: targetEntity,
		Description:  "Builds a solar grid on the target Planet",
		ExecuteTick:  executeTick,
		Duration:     duration, // TODO: eventually Tick * 40 for clarity
		Status:       Pending,
		Order:        order,
	}

	return action
}

func (a BuildSolarGrid) GetID() int                   { return a.ID }
func (a BuildSolarGrid) GetTargetEntity() Entity      { return a.TargetEntity }
func (a BuildSolarGrid) GetDescription() string       { return a.Description }
func (a BuildSolarGrid) GetExecuteTick() int          { return a.ExecuteTick }
func (a BuildSolarGrid) GetDuration() int             { return a.Duration }
func (a BuildSolarGrid) GetStatus() EventStatus       { return a.Status }
func (a BuildSolarGrid) SetStatus(status EventStatus) { a.Status = status }
func (a BuildSolarGrid) GetOrder() Order              { return a.Order }

func (a *BuildSolarGrid) Execute() {
	a.Status = Executing

	if a.ExecuteTick+a.Duration <= State.Tick {
		SolarGrid := constructions.CreateSolarGrid(1)

		if planetEntity, ok := a.TargetEntity.(*Planet); ok {
			planetEntity.Constructions.SolarGrids = append(planetEntity.Constructions.SolarGrids, SolarGrid)
			logging.Info("%v: Added SolarGrid", planetEntity.GetName())
		}

		a.Status = Complete
	}
}
