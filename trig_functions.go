package main

import (
	"math"
)

func sin(deg float64) float64 {
	return math.Sin(math.Radians(deg))
}

func asin(x float64) float64 {
	return math.Asin(x) * 180 / math.Pi
}

func tan(deg float64) float64 {
	return math.Tan(math.Radians(deg))
}

func cos(deg float64) float64 {
	return math.Cos(math.Radians(deg))
}
