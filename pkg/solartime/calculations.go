package solartime

import (
    "math"
    "time"
)

// equationOfTime calculates the difference between apparent solar time and mean solar time.
// Formula from: https://en.wikipedia.org/wiki/Equation_of_time
func equationOfTime(t time.Time) float64 {
    year := float64(t.Year())
    dayOfYear := float64(t.YearDay())

    D := 6.24004077 + 0.01720197*(365.25*(year-2000)+dayOfYear)
    eot := -7.659*math.Sin(D) + 9.863*math.Sin(2*D+3.5932)

    return eot
}
