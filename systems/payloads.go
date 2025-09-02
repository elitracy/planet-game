package systems

import (
	// "fmt"

	"github.com/elitracy/planets/models"
)

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

			// ready_message_payloads := p.ReadMessagePayloads(gs.CurrentTick)
			// ready_resource_payloads := p.ReadResourcePayloads(gs.CurrentTick)

			// update for player input
			// fmt.Printf("%s: Message Payloads:\n%v\n", p.Name, ready_message_payloads)
			// fmt.Printf("%s: Resource Payloads:\n%v\n", p.Name, ready_resource_payloads)
		}

	}

}
