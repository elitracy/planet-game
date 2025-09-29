package systems

import (
	"github.com/elitracy/planets/logging"
	. "github.com/elitracy/planets/models"
)

func TickPlanets(state *GameState) {
	for _, system := range state.StarSystems {
		for _, planet := range system.Planets {

			for _, order := range planet.GetOrders() {
				switch order.Status {
				case Pending:
					if order.StartTime == state.CurrentTick {
						logging.Info("[PLANET %v] Executing Order: %v", planet.Name, order.ID)
						order.Status = Executing
						for _, a := range order.Actions {
							a.Execute()
						}
					}
				case Executing:
					complete := true
					for _, a := range order.Actions {
						if a.Status != Complete {
							complete = false
						}

						if a.Status == Failed {
							order.Status = Failed
						}
					}

					if complete {
						order.Status = Complete
					}

				case Complete:
					order := planet.PopOrder()
					logging.Info("[PLANET %v] Completed Order: %v", planet.Name, order.ID)
				case Failed:
					logging.Error("[PLANET %v] Order Failed: %v", planet.Name, order.ID)
				}
			}
		}
	}

}
