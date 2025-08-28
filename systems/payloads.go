package systems

import "github.com/elitracy/planets/models"

func TickPayloads(gs *models.GameState) {
	for _, s := range gs.StarSystems {
		for _, p := range s.Planets {
			for _, payload := range p.MessagePayloads {
				if payload.TimeArrival == gs.CurrentTick {
					payload.Arrived = true
				}

			}
			for _, payload := range p.ResourcePayloads {
				if payload.TimeArrival == gs.CurrentTick {
					payload.Arrived = true
				}

			}
		}

	}

}
