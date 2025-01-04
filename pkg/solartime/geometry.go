package solartime

import "math"

func geomMeanLongSun(jme float64) float64 {
	return 280.46646 + jme*(36000.76983+jme*0.0003032) - 360*math.Floor(jme*(36000.76983+jme*0.0003032)/360)
}

func geomMeanAnomalySun(jme float64) float64 {
	return 357.52911 + jme*(35999.05029-0.0001537*jme)
}
