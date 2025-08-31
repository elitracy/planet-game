package models

import (
	"github.com/elitracy/planets/models/resources"
)

const PLAYER_PAYLOAD_SPEED = 10

type Player struct {
	Location
}

// only send to planets for now
func (p *Player) SendMessagePayload(msg string, planet *Planet, currentTick int) {

	distance := Distance(p.Location, planet.Location)

	payload := Payload[string]{
		Data:        msg,
		Origin:      p.Location,
		Destination: planet.Location,
		TimeSent:    currentTick,
		TimeArrival: int(distance / PLAYER_PAYLOAD_SPEED),
		Arrived:     false,
	}

	planet.PlanetPayloads.MessagePayloads = append(planet.PlanetPayloads.MessagePayloads, &payload)
}

// only send to planets for now
func (p *Player) SendResourcePayload(resource resources.Resource, planet *Planet, currentTick int) {

	distance := Distance(p.Location, planet.Location)

	payload := Payload[resources.Resource]{
		Data:        resource,
		Origin:      p.Location,
		Destination: planet.Location,
		TimeSent:    currentTick,
		TimeArrival: int(distance / PLAYER_PAYLOAD_SPEED),
		Arrived:     false,
	}

	planet.PlanetPayloads.ResourcePayloads = append(planet.PlanetPayloads.ResourcePayloads, &payload)
}
