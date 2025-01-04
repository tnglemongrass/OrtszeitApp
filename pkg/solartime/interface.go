package solartime

import "time"

// Location represents a geographical position
type Location struct {
    Latitude  float64
    Longitude float64
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
