package systems

import (
	"github.com/elitracy/planets/core/consts"
	"github.com/elitracy/planets/core/logging"
	"github.com/elitracy/planets/models/events/orders"
	"github.com/elitracy/planets/state"
)

func TickOrderScheduler() {
	for _, order := range state.State.OrderScheduler.PriorityQueue {
		switch order.GetStatus() {
		case consts.EventPending:
			if order.GetExecuteTick() <= state.State.Tick {
				logging.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(consts.EventExecuting)
				for _, action := range order.GetActions() {
					if action.GetExecuteTick() == state.State.Tick {
						action.SetStatus(consts.EventExecuting)
					}
				}
			}
		case consts.EventExecuting:
			complete := true
			for _, action := range order.GetActions() {
				if action.GetStatus() != consts.EventComplete {
					complete = false
				}

				if action.GetStatus() == consts.Failed {
					order.SetStatus(consts.Failed)
				}
			}

			if complete {
				order.SetStatus(consts.EventComplete)
			}

		case consts.EventComplete: // should be safe to pop order scheulder queue here
			poppedOrder := state.State.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				logging.Error("Order Scheduler out of sync")
				logging.Error("Order Sheduler Queue: %v", state.State.OrderScheduler)
				logging.Error("Popped Order: %v", order)
				logging.Error("Expected Order: %v", poppedOrder)
			}

			state.State.CompletedOrders = append([]*orders.Order{poppedOrder}, state.State.CompletedOrders...)

			logging.Info("[%v] Completed Order: %v", order.GetName(), order.GetID())

		case consts.Failed:
			logging.Error("[%v] Order Failed: %v", order.GetName(), order.GetID())
		}
	}
}

func TickActionScheduler() {
	for _, order := range state.State.OrderScheduler.PriorityQueue {
		if order.GetStatus() == consts.EventExecuting {
			for _, action := range order.GetActions() {

				switch action.GetStatus() {
				case consts.EventPending:
					if action.GetExecuteTick() <= state.State.Tick {
						action.SetStatus(consts.EventExecuting)
					}
				case consts.EventExecuting:
					if action.GetExecuteTick()+action.GetDuration() <= state.State.Tick {
						action.Execute()
						action.SetStatus(consts.EventComplete)
					}
				}

			}
		}
	}
}
