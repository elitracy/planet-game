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
			if order.GetStartTick() <= state.State.CurrentTick {
				logging.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(events.EventExecuting)
				for _, action := range order.GetActions() {
					if action.GetStartTick() == state.State.CurrentTick {
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
	for _, action := range state.State.ActionScheduler.PriorityQueue {
		switch action.GetStatus() {
		case events.EventPending:
			if action.GetStartTick() <= state.State.CurrentTick {
				action.SetStatus(events.EventExecuting)
			}
		case events.EventExecuting:
			if action.GetStartTick()+action.GetDuration() <= state.State.CurrentTick {
				action.Execute()
				action.SetStatus(events.EventComplete)
			}
		case events.EventComplete:
			poppedAction := state.State.ActionScheduler.Pop()
			if poppedAction.GetID() != action.GetID() {
				logging.Error("Action Scheduler out of sync")
				logging.Error("Popped Action: %v", action)
				logging.Error("Expected Action: %v", poppedAction)
			}

			logging.Info("[%v] Completed Action: %v", action.GetDescription(), action.GetID())

		case events.EventFailed:
			logging.Error("[%v] Action Failed: %v", action.GetDescription(), action.GetID())

		}
	}
}
