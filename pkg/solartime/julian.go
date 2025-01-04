package solartime

import "time"

func julianDay(t time.Time) float64 {
	y, m, d := t.Date()
	h := float64(t.Hour())
	min := float64(t.Minute())
	sec := float64(t.Second())
	return 367*float64(y) - float64(7*(y+int((m+9)/12))/4) + float64(275*m/9) + float64(d) + 1721013.5 + (h+min/60+sec/3600)/24.0
}

func julianCentury(jd float64) float64 {
	return (jd - 2451545.0) / 36525.0
}

func julianEphemerisDay(jc float64) float64 {
	return 2451545.0 + jc*36525.0
}

func julianEphemerisCentury(jde float64) float64 {
	return (jde - 2451545.0) / 36525.0
}

func julianEphemerisMillennium(jce float64) float64 {
	return jce / 10.0
}
