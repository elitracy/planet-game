package models

import (
	"fmt"
	"math"
)

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) String() string {
	var output string
	output += fmt.Sprintf("(%d,%d)", c.X, c.Y)
	return output
}

type Location struct {
	Coordinates Coordinates
}

func (l *Location) String() string {
	return l.Coordinates.String()
}

// l2 - l1 btw
func Distance(l1 Location, l2 Location) float64 {
	euclidian_distance := math.Sqrt(
		math.Pow(float64(l2.Coordinates.X-l1.Coordinates.X), 2) + math.Pow(float64(l2.Coordinates.Y-l1.Coordinates.Y), 2),
	)

	return euclidian_distance

}
