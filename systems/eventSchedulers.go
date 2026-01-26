package systems

import (
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models/events"
	"github.com/elitracy/planets/models/events/orders"
	"github.com/elitracy/planets/state"
)

func TickOrderScheduler() {
	for _, order := range state.State.OrderScheduler.PriorityQueue {
		switch order.GetStatus() {
		case events.EventPending:
			if order.GetExecuteTick() <= state.State.CurrentTick {
				logging.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(events.EventExecuting)
				for _, action := range order.GetActions() {
					if action.GetExecuteTick() == state.State.CurrentTick {
						action.SetStatus(events.EventExecuting)
					}
				}
			}
		case events.EventExecuting:
			complete := true
			for _, action := range order.GetActions() {
				if action.GetStatus() != events.EventComplete {
					complete = false
				}

				if action.GetStatus() == events.EventFailed {
					order.SetStatus(events.EventFailed)
				}
			}

			if complete {
				order.SetStatus(events.EventComplete)
			}

		case events.EventComplete: 
			poppedOrder := state.State.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				logging.Error("Order Scheduler out of sync")
				logging.Error("Order Sheduler Queue: %v", state.State.OrderScheduler)
				logging.Error("Popped Order: %v", order)
				logging.Error("Expected Order: %v", poppedOrder)
			}

			state.State.CompletedOrders = append([]*orders.Order{poppedOrder}, state.State.CompletedOrders...)

			logging.Info("[%v] Completed Order: %v", order.GetName(), order.GetID())

		case events.EventFailed:
			logging.Error("[%v] Order Failed: %v", order.GetName(), order.GetID())
		}
	}
}

func TickActionScheduler() {
	for _, order := range state.State.OrderScheduler.PriorityQueue {
		if order.GetStatus() == events.EventExecuting {
			for _, action := range order.GetActions() {

				switch action.GetStatus() {
				case events.EventPending:
					if action.GetExecuteTick() <= state.State.CurrentTick {
						action.SetStatus(events.EventExecuting)
					}
				case events.EventExecuting:
					if action.GetExecuteTick()+action.GetDuration() <= state.State.CurrentTick {
						action.Execute()
						action.SetStatus(events.EventComplete)
					}
				}

			}
		}
	}
}
