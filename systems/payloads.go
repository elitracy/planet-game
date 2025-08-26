package systems

import "github.com/elitracy/planets/models"

func TickPayloads(gs *models.GameState) {
	for _, p := range gs.MessagePayLoads {

		// TODO: add message to [destination queue]
		if p.TimeArrival == gs.CurrentTick {
			p.Arrived = true
		}
	}

	for _, p := range gs.ResourcePayLoads {

		// TODO: add resources to destination
		if p.TimeArrival == gs.CurrentTick {
			p.Arrived = true
		}
	}
}
