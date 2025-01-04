package main

import (
	"math"
)

func eclipticLongSun(m, l0 float64) float64 {
	return l0 + 1.9148*sin(m) + 0.02*sin(2*m)
}

func eccentricityEarthOrbit(jme float64) float64 {
	return 0.016708634 - jme*(0.000042037+0.0000001267*jme)
}

func sunEqOfCenter(m, l float64) float64 {
	return 1.9148*sin(m) + 0.02*sin(2*m) + 0.0003*sin(3*m)
}

func trueSolarDeclination(l float64) float64 {
	return asin(sin(l) * sin(23.44))
}

func equationOfTime(m, l, e float64) float64 {
	l1 := eclipticLongSun(m, l)
	y := tan(2 * asin(sin(l1) * sin(23.44)))
	return y*sin(2*l) - 2*e*sin(m) + 4*e*y*sin(m)*cos(2*l) - 0.5*y*y*sin(4*l)
}

func trueSolarTime(t time.Time, longitude, eot float64) time.Time {
	lst := 100.46 + 0.985647*float64(t.Hour()) + longitude + eot
	return t.Add(time.Duration(lst-12) * time.Hour / 15)
}

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
