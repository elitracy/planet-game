package systems

import (
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
	. "github.com/elitracy/planets/core/state"
	. "github.com/elitracy/planets/models"
)

func TickOrderScheduler() {
	for _, order := range State.OrderScheduler.PriorityQueue {
		switch order.GetStatus() {
		case consts.Pending:
			if order.GetExecuteTick() <= State.Tick {
				logging.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(consts.Executing)
				for _, action := range order.GetActions() {
					if action.GetExecuteTick() == State.Tick {
						action.SetStatus(consts.Executing)
					}
				}
			}
		case consts.Executing:
			complete := true
			for _, action := range order.GetActions() {
				if action.GetStatus() != consts.Complete {
					complete = false
				}

				if action.GetStatus() == consts.Failed {
					order.SetStatus(consts.Failed)
				}
			}

			if complete {
				order.SetStatus(consts.Complete)
			}

		case consts.Complete: // should be safe to pop order scheulder queue here
			poppedOrder := State.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				logging.Error("Order Scheduler out of sync")
				logging.Error("Order Sheduler Queue: %v", State.OrderScheduler)
				logging.Error("Popped Order: %v", order)
				logging.Error("Expected Order: %v", poppedOrder)
			}

			State.CompletedOrders = append([]Order{poppedOrder}, State.CompletedOrders...)

			logging.Info("[%v] Completed Order: %v", order.GetName(), order.GetID())

		case consts.Failed:
			logging.Error("[%v] Order Failed: %v", order.GetName(), order.GetID())
		}
	}
}

func TickActionScheduler() {
	for _, order := range State.OrderScheduler.PriorityQueue {
		if order.GetStatus() == consts.Executing {
			for _, action := range order.GetActions() {

				switch action.GetStatus() {
				case consts.Pending:
					if action.GetExecuteTick() <= State.Tick {
						action.SetStatus(consts.Executing)
					}
				case consts.Executing:
					if action.GetExecuteTick()+action.GetDuration() <= State.Tick {
						action.Execute()
						action.SetStatus(consts.Complete)
					}
				}

			}
		}
	}
}
