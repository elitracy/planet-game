package models

import "fmt"

type Payload[T any] struct {
	Data        T
	Origin      Location
	Destination Location
	TimeSent    int
	TimeArrival int
	Arrived     bool
}

func (p *Payload[T]) String() string {

	var output string

	output += fmt.Sprintf("> [%v]\n", p.Data)
	output += fmt.Sprintf("| Origin:       %v\n", p.Origin)
	output += fmt.Sprintf("| Dest:         %v\n", p.Destination)
	output += fmt.Sprintf("| Tick Sent:    %v\n", p.TimeSent)
	output += fmt.Sprintf("| Tick Arrival: %v\n", p.TimeArrival)
	output += fmt.Sprintf("| Arrived:      %v", p.Arrived)
	return output
}
