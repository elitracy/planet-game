package game

import (
	"github.com/elitracy/planets/engine"
	"github.com/elitracy/planets/engine/task"
	"github.com/elitracy/planets/game/orders"
)

func (state *GameState) TickOrders() {
	for _, order := range state.OrderScheduler.PriorityQueue {

		switch order.GetStatus() {
		case task.Pending:
			if order.GetStartTick() <= state.CurrentTick {
				engine.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(task.Executing)
				for _, action := range order.GetActions() {
					if action.GetStartTick() == state.CurrentTick {
						action.SetStatus(task.Executing)
					}
				}
			}
		case task.Executing:
			complete := true
			for _, action := range order.GetActions() {
				if action.GetStatus() != task.Complete {
					complete = false
				}

				if action.GetStatus() == task.Failed {
					order.SetStatus(task.Failed)
				}
			}

			if complete {
				order.SetStatus(task.Complete)
			}

		case task.Complete:
			poppedOrder := state.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				engine.Error("Order Scheduler out of sync")
				engine.Error("Popped Order: %v", order)
				engine.Error("Expected Order: %v", poppedOrder)
			}

			state.CompletedOrders = append([]*orders.Order{poppedOrder}, state.CompletedOrders...)

			engine.Info("[%v] Completed Order: %v", order.GetName(), order.GetID())

		case task.Failed:
			engine.Error("[%v] Order Failed: %v", order.GetName(), order.GetID())
		}
	}
}

func (state *GameState) TickActions() {
	for _, action := range state.ActionScheduler.PriorityQueue {
		switch action.GetStatus() {
		case task.Pending:
			if action.GetStartTick() <= state.CurrentTick {
				action.SetStatus(task.Executing)
			}
		case task.Executing:
			if action.GetStartTick()+action.GetDuration() <= state.CurrentTick {
				if action.Execute != nil {
					action.Execute()
				}
				action.SetStatus(task.Complete)
			}
		case task.Complete:
			poppedAction := state.ActionScheduler.Pop()
			if poppedAction.GetID() != action.GetID() {
				engine.Error("Action Scheduler out of sync")
				engine.Error("Popped Action: %v", action)
				engine.Error("Expected Action: %v", poppedAction)
			}

			engine.Info("[%v] Completed Action: %v", action.GetDescription(), action.GetID())

		case task.Failed:
			engine.Error("[%v] Action Failed: %v", action.GetDescription(), action.GetID())

		}
	}
}
