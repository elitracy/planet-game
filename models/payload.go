package models

type Payload[T any] struct {
	Data        T
	Origin      Location
	Destination Location
	TimeSent    int
	TimeArrival int
	Arrived     bool
}
