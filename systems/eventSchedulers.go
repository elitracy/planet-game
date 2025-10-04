package systems

import (
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
)

func TickOrderScheduler(gs *GameState) {

	for _, order := range gs.OrderScheduler.PriorityQueue {
		switch order.Status {
		case Pending:
			if order.StartTime == gs.CurrentTick {
				logging.Info("[%s] Executing Order: %v", (*order.TargetEntity).GetName(), order.ID)
				order.Status = Executing
				for _, a := range order.Actions {
					a.Status = Executing
				}
			}
		case Executing:
			complete := true
			for _, a := range order.Actions {
				if a.Status != Complete {
					complete = false

					// order.Tick() // events should have a tick/update function
				}

				if a.Status == Failed {
					order.Status = Failed
				}
			}

			if complete {
				order.Status = Complete
			}

		case Complete: // should be safe to pop order scheulder queue here
			poppedOrder := gs.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				logging.Error("Order Scheduler out of sync")
				logging.Error("Order Sheduler Queue: %v", gs.OrderScheduler)
				logging.Error("Popped Order: %v", *order)
				logging.Error("Expected Order: %v", *poppedOrder)
			}
			logging.Info("[%v] Completed Order: %v", (*order.TargetEntity).GetName(), order.ID)
		case Failed:
			logging.Error("[%v] Order Failed: %v", (*order.TargetEntity).GetName(), order.ID)
		}
	}
}
