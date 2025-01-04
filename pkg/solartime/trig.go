package solartime

import "math"

const degToRad = math.Pi / 180.0
const radToDeg = 180.0 / math.Pi

func sin(deg float64) float64 {
	return math.Sin(deg * degToRad)
}

func asin(x float64) float64 {
	return math.Asin(x) * radToDeg
}

func tan(deg float64) float64 {
	return math.Tan(deg * degToRad)
}

func cos(deg float64) float64 {
	return math.Cos(deg * degToRad)
}
