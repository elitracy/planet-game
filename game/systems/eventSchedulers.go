package systems

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/game"
	"github.com/elitracy/planets/game/orders"
)

func TickOrderScheduler() {
	for _, order := range game.State.OrderScheduler.PriorityQueue {

		switch order.GetStatus() {
		case engine.EventPending:
			if order.GetStartTick() <= game.State.CurrentTick {
				engine.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(engine.EventExecuting)
				for _, action := range order.GetActions() {
					if action.GetStartTick() == game.State.CurrentTick {
						action.SetStatus(engine.EventExecuting)
					}
				}
			}
		case engine.EventExecuting:
			complete := true
			for _, action := range order.GetActions() {
				if action.GetStatus() != engine.EventComplete {
					complete = false
				}

				if action.GetStatus() == engine.EventFailed {
					order.SetStatus(engine.EventFailed)
				}
			}

			if complete {
				order.SetStatus(engine.EventComplete)
			}

		case engine.EventComplete:
			poppedOrder := game.State.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				engine.Error("Order Scheduler out of sync")
				engine.Error("Popped Order: %v", order)
				engine.Error("Expected Order: %v", poppedOrder)
			}

			game.State.CompletedOrders = append([]*orders.Order{poppedOrder}, game.State.CompletedOrders...)

			engine.Info("[%v] Completed Order: %v", order.GetName(), order.GetID())

		case engine.EventFailed:
			engine.Error("[%v] Order Failed: %v", order.GetName(), order.GetID())
		}
	}
}

func TickActionScheduler() {
	for _, action := range game.State.ActionScheduler.PriorityQueue {
		switch action.GetStatus() {
		case engine.EventPending:
			if action.GetStartTick() <= game.State.CurrentTick {
				action.SetStatus(engine.EventExecuting)
			}
		case engine.EventExecuting:
			if action.GetStartTick()+action.GetDuration() <= game.State.CurrentTick {
				action.Execute()
				action.SetStatus(engine.EventComplete)
			}
		case engine.EventComplete:
			poppedAction := game.State.ActionScheduler.Pop()
			if poppedAction.GetID() != action.GetID() {
				engine.Error("Action Scheduler out of sync")
				engine.Error("Popped Action: %v", action)
				engine.Error("Expected Action: %v", poppedAction)
			}

			engine.Info("[%v] Completed Action: %v", action.GetDescription(), action.GetID())

		case engine.EventFailed:
			engine.Error("[%v] Action Failed: %v", action.GetDescription(), action.GetID())

		}
	}
}
