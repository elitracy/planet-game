package orders

import (
	"github.com/elitracy/planets/core"
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
	. "github.com/elitracy/planets/core/state"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/actions"
)

type SendScoutShipOrder struct {
	ID          int
	Name        string
	Actions     []Action
	ExecuteTick core.Tick
	Status      consts.EventStatus
	*Ship
	Destination core.Position
}

func NewScoutShipOrder(ship *Ship, dest core.Position, execTick core.Tick) *SendScoutShipOrder {
	order := &SendScoutShipOrder{
		ID:          State.OrderScheduler.GetNextID(),
		Name:        "Send Scout Ship",
		ExecuteTick: execTick,
		Status:      consts.Pending,
		Ship:        ship,
		Destination: dest,
	}

	d := core.EuclidianDistance(ship.Position, dest)
	t := d / ship.Velocity.Vector()
	logging.Info("PLANET POS: %v", dest)
	logging.Info("SHIP POS:   %v", ship.Position)
	logging.Info("DISTANCE:   %v", d)
	logging.Info("TIME (s):   %v", t/TICKS_PER_SECOND)

	travelAction := actions.NewMoveShipAction(
		ship,
		order.ExecuteTick,
		core.Tick(t),
		dest,
	)

	order.Actions = append(order.Actions, travelAction)

	return order

}

func (o SendScoutShipOrder) GetID() int                    { return o.ID }
func (o SendScoutShipOrder) GetName() string               { return o.Name }
func (o SendScoutShipOrder) GetActions() []Action          { return o.Actions }
func (o SendScoutShipOrder) GetExecuteTick() core.Tick     { return o.ExecuteTick }
func (o SendScoutShipOrder) GetStatus() consts.EventStatus { return o.Status }

func (o SendScoutShipOrder) GetDuration() core.Tick {
	var duration core.Tick
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}

func (o *SendScoutShipOrder) SetStatus(status consts.EventStatus) { o.Status = status }
