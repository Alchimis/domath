package geom

import "math"

func Distance2D(a, b Point2D) float64 {
	return math.Sqrt(math.Pow((a.X-b.X), 2) + math.Pow(a.Y-b.Y, 2))
}
