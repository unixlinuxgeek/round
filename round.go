package round

import (
	"math"
	"math/rand"
)

// round to the nearest
func Round(x float64) float64 {
	return math.Round(x*100) / 100
}

// round up
func Ceil(x float64) float64 {
	return math.Ceil(x*100) / 100
}

// round down
func Floor(x float64) float64 {
	return math.Floor(x*100) / 100
}

func genRandomFloat(x, y float64) float64 {
	return 1 + rand.Float64()*(x-y) // generate random float
}
