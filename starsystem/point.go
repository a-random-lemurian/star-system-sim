package starsystem

import "math"

type Point struct {
	x float64
	y float64
}

func Distance(p1, p2 Point) float64 {
	// Convert int64 to float64 for the distance calculation
	dx := float64(p2.x - p1.x)
	dy := float64(p2.y - p1.y)
	return math.Sqrt(dx*dx + dy*dy)
}
