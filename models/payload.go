package models

type Payload[T any] struct {
	Data        T
	Origin      Coordinates
	Destination Coordinates
	TimeSent    int
	TimeArrival int
}
