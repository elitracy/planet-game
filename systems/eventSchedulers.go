package systems

import (
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
	. "github.com/elitracy/planets/state"
)

func TickOrderScheduler() {
	for _, order := range State.OrderScheduler.PriorityQueue {
		switch order.GetStatus() {
		case Pending:
			if order.GetExecuteTick() <= State.Tick {
				logging.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(Executing)
				for _, action := range order.GetActions() {
					if action.GetExecuteTick() == State.Tick {
						action.SetStatus(Executing)
					}
				}
			}
		case Executing:
			complete := true
			for _, action := range order.GetActions() {
				if action.GetStatus() != Complete {
					complete = false
				}

				if action.GetStatus() == Failed {
					order.SetStatus(Failed)
				}
			}

			if complete {
				order.SetStatus(Complete)
			}

		case Complete: // should be safe to pop order scheulder queue here
			poppedOrder := State.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				logging.Error("Order Scheduler out of sync")
				logging.Error("Order Sheduler Queue: %v", State.OrderScheduler)
				logging.Error("Popped Order: %v", order)
				logging.Error("Expected Order: %v", poppedOrder)
			}

			State.CompletedOrders = append([]Order{poppedOrder}, State.CompletedOrders...)

			logging.Info("[%v] Completed Order: %v", order.GetName(), order.GetID())

		case Failed:
			logging.Error("[%v] Order Failed: %v", order.GetName(), order.GetID())
		}
	}
}

func TickActionScheduler() {
	for _, order := range State.OrderScheduler.PriorityQueue {
		if order.GetStatus() == Executing {
			for _, action := range order.GetActions() {

				if action.GetExecuteTick() <= State.Tick {
					action.Execute()
				}
			}
		}
	}
}
