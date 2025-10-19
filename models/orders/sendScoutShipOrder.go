package orders

import (
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
	"github.com/elitracy/planets/models/actions"
	. "github.com/elitracy/planets/state"
)

type SendScoutShipOrder struct {
	ID          int
	Name        string
	Actions     []Action
	ExecuteTick int
	Status      EventStatus
	*Ship
	Destination Position
}

func NewScoutShipOrder(ship *Ship, dest Position, execTick int) *SendScoutShipOrder {
	order := &SendScoutShipOrder{
		ID:          State.OrderScheduler.GetNextID(),
		Name:        "Send Scout Ship",
		ExecuteTick: execTick,
		Status:      Pending,
		Ship:        ship,
		Destination: dest,
	}

	d := EuclidianDistance(ship.Position, dest)
	t := d / ship.Velocity.Vector()
	logging.Info("PLANET POS: %v", dest)
	logging.Info("SHIP POS:   %v", ship.Position)
	logging.Info("DISTANCE:   %v", d)
	logging.Info("TIME (s):   %v", t/TICKS_PER_SECOND)

	travelAction := actions.NewMoveShipAction(
		ship,
		order.ExecuteTick,
		int(t),
		dest,
	)

	order.Actions = append(order.Actions, travelAction)

	return order

}

func (o SendScoutShipOrder) GetID() int             { return o.ID }
func (o SendScoutShipOrder) GetName() string        { return o.Name }
func (o SendScoutShipOrder) GetActions() []Action   { return o.Actions }
func (o SendScoutShipOrder) GetExecuteTick() int    { return o.ExecuteTick }
func (o SendScoutShipOrder) GetStatus() EventStatus { return o.Status }

func (o SendScoutShipOrder) GetDuration() int {
	duration := 0
	for _, a := range o.Actions {
		duration += a.GetDuration()
	}
	return duration
}

func (o *SendScoutShipOrder) SetStatus(status EventStatus) { o.Status = status }
