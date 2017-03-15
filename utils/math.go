package utils

import (
	"math"

	"engo.io/engo"
)

func PointsDirection(p1 engo.Point, p2 engo.Point) float64 {
	d := math.Atan2(float64(p2.Y-p1.Y), float64(p2.X-p1.X)) * 180 / math.Pi
	if d < 0 {
		return d + 360
	}

	return d
}
