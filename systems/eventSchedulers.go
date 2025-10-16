package systems

import (
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
)

func TickOrderScheduler() {
	for _, order := range GameStateGlobal.OrderScheduler.PriorityQueue {
		switch order.GetStatus() {
		case Pending:
			logging.Info("Order exec tick: %v", order.GetExecuteTick())
			if order.GetExecuteTick() <= GameStateGlobal.CurrentTick {
				logging.Info("[%s] Executing Order: %v", order.GetName(), order.GetID())
				order.SetStatus(Executing)
				for _, action := range order.GetActions() {
					if action.ExecuteTick == GameStateGlobal.CurrentTick {
						action.Status = Executing
					}
				}
			}
		case Executing:
			complete := true
			for _, action := range order.GetActions() {
				if action.Status != Complete {
					complete = false
				}

				if action.Status == Failed {
					order.SetStatus(Failed)
				}
			}

			if complete {
				order.SetStatus(Complete)
			}

		case Complete: // should be safe to pop order scheulder queue here
			poppedOrder := GameStateGlobal.OrderScheduler.Pop()
			if poppedOrder.GetID() != order.GetID() {
				logging.Error("Order Scheduler out of sync")
				logging.Error("Order Sheduler Queue: %v", GameStateGlobal.OrderScheduler)
				logging.Error("Popped Order: %v", order)
				logging.Error("Expected Order: %v", poppedOrder)
			}

			GameStateGlobal.CompletedOrders = append([]Order{poppedOrder}, GameStateGlobal.CompletedOrders...)

			logging.Info("[%v] Completed Order: %v", order.GetName(), order.GetID())

		case Failed:
			logging.Error("[%v] Order Failed: %v", order.GetName(), order.GetID())
		}
	}
}

func TickActionScheduler() {
	for _, order := range GameStateGlobal.OrderScheduler.PriorityQueue {
		if order.GetStatus() == Executing {
			for _, action := range order.GetActions() {
				if action.ExecuteTick+action.Duration == GameStateGlobal.CurrentTick {
					action.Execute()
				}
			}
		}
	}
}
