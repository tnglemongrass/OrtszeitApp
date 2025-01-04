package solartime

import (
	"math"
	"time"
)

// Location represents a geographical location
type Location struct {
	Latitude  float64
	Longitude float64
}

// Calculate returns the solar time for a given location
func Calculate(loc Location) string {
	now := time.Now()
	return CalculateAt(loc, now)
}

// CalculateAt returns the solar time for a given location and time
func CalculateAt(loc Location, t time.Time) string {
	solarTime := GetSolarTime(loc, t)
	return solarTime.Format("15:04:05")
}

// GetSolarTime calculates the real solar time for a given location
func GetSolarTime(loc Location, t time.Time) time.Time {
	// Convert longitude to hour offset (15° = 1 hour)
	longitudeHourOffset := loc.Longitude / 15.0

	// Get the equation of time correction in minutes
	eot := equationOfTime(t)

	// Convert to minutes and add longitude correction
	totalOffsetMinutes := eot + (longitudeHourOffset * 60)

	// Apply the offset to the current time
	return t.Add(time.Duration(totalOffsetMinutes) * time.Minute)
}

// equationOfTime calculates the difference between apparent solar time and mean solar time.
// Formula from: https://en.wikipedia.org/wiki/Equation_of_time
func equationOfTime(t time.Time) float64 {
	year := float64(t.Year())
	dayOfYear := float64(t.YearDay())

	D := 6.24004077 + 0.01720197*(365.25*(year-2000)+dayOfYear)
	eot := -7.659*math.Sin(D) + 9.863*math.Sin(2*D+3.5932)

	return eot
}
