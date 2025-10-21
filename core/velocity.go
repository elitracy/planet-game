package core

import "math"

type Velocity struct {
	X int
	Y int
	Z int
}

func (v Velocity) Vector() float64 {
	return math.Sqrt(
		math.Pow(float64(v.X), 2.0) +
			math.Pow(float64(v.Y), 2.0) +
			math.Pow(float64(v.Z), 2.0),
	)
}
