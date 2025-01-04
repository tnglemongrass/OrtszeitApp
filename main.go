package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func calculateSolarTime(latitude, longitude float64) string {
	now := time.Now()
	jd := julianDay(now)
	jc := julianCentury(jd)
	jde := julianEphemerisDay(jc)
	jce := julianEphemerisCentury(jde)
	jme := julianEphemerisMillennium(jce)
	l0 := geomMeanLongSun(jme)
	m := geomMeanAnomalySun(jme)
	l := eclipticLongSun(m, l0)
	e := eccentricityEarthOrbit(jme)
	eot := equationOfTime(m, l, e)
	tst := trueSolarTime(now, longitude, eot)
	return tst.Format("15:04:05")
}

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

func geomMeanLongSun(jme float64) float64 {
	return 280.46646 + jme*(36000.76983+jme*0.0003032) - 360*math.Floor(jme*(36000.76983+jme*0.0003032)/360)
}

func geomMeanAnomalySun(jme float64) float64 {
	return 357.52911 + jme*(35999.05029-0.0001537*jme)
}

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
	y := tan(2 * asin(sin(l1)*sin(23.44)))
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

func main() {
	a := app.New()
	w := a.NewWindow("Ortszeit")

	// Create a label to display the Ortszeit
	ortszeitLabel := widget.NewLabel("Fetching Ortszeit...")

	// Set the content of the window
	w.SetContent(container.NewVBox(
		ortszeitLabel,
	))

	// Placeholder for location fetching
	// For demonstration, we will use a fixed location (e.g., latitude: 37.7749, longitude: -122.4194)
	latitude := 37.7749
	longitude := -122.4194

	ortszeit := calculateSolarTime(latitude, longitude)
	ortszeitLabel.SetText(ortszeit)

	// Show the window and run the app
	w.ShowAndRun()
}
