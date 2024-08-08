package utils

import "math"

type Vector struct {
	X, Y int
}

func Distance(a Vector, b Vector) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}
