package coords

import (
	"fmt"
	"math"
)

type Coord2D struct {
	x float64
	y float64
}

func getCoord() Coord2D {
	return Coord2D{x: 1, y: 2}
}

func Distance2D(p1, p2 Coord2D) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

// TODO: IMPLEMENT
func FindMinDist2D(coords ...[]Coord2D) {
	if len(coords) == 1 {
		for _, c := range coords[0] {
			fmt.Println(c.x)
		}
	}
}
