package core

import (
	"fmt"
	"math"

	"github.com/dustin/go-humanize"
)

type Position struct {
	X int64
	Y int64
	Z int64
}

func (l Position) String() string {
	return fmt.Sprintf("(%v, %v, %v)", humanize.Comma(l.X), humanize.Comma(l.Y), humanize.Comma(l.Z))
}

func EuclidianDistance(l1 Position, l2 Position) float64 {
	euclidian_distance := math.Sqrt(
		math.Pow(float64(l2.X-l1.X), 2) + math.Pow(float64(l2.Y-l1.Y), 2) + math.Pow(float64(l2.Z-l1.Z), 2),
	)

	return euclidian_distance
}

func (p0 Position) Add(p1 Position) Position {
	return Position{p0.X + p1.X, p0.Y + p1.Y, p0.Z + p1.Z}
}

func (p0 Position) Sub(p1 Position) Position {
	return Position{p0.X - p1.X, p0.Y - p1.Y, p0.Z - p1.Z}
}
