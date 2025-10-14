package models

import (
	"fmt"
	"math"
)

type Position struct {
	X int
	Y int
	Z int
}

func (l Position) String() string {
	return fmt.Sprintf("(%d, %d, %d)", l.X, l.Y, l.Z)
}

func EuclidianDistance(l1 Position, l2 Position) float64 {
	euclidian_distance := math.Sqrt(
		math.Pow(float64(l2.X-l1.X), 2) + math.Pow(float64(l2.Y-l1.Y), 2) + math.Pow(float64(l2.Z-l1.Z), 2),
	)

	return euclidian_distance
}
