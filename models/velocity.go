package models

import "math"

type Velocity struct {
	X int
	Y int
	Z int
}

func (v Velocity) Vector() int {
	return int(math.Sqrt(
		math.Pow(float64(v.X), 2.0) +
			math.Pow(float64(v.Y), 2.0) +
			math.Pow(float64(v.Z), 2.0),
	))
}
