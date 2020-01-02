package src

import "math"

type Point struct{ X, Y float64 }
type Path []Point

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func Distance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
			//dd(i,path[i-1],path[i],sum)
		}
	}
	return sum
}
