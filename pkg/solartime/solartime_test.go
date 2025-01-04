package solartime

import (
    "testing"
    "time"
)

func TestGetSolarTime(t *testing.T) {
    // Test location (Munich)
    loc := Location{
        Latitude:  48.1351,
        Longitude: 11.5820,
    }

    // Test time
    testTime := time.Date(2024, 3, 20, 12, 0, 0, 0, time.UTC)
    
    // Get solar time
    solarTime := GetSolarTime(loc, testTime)

    // Verify we get a valid time back
    if solarTime.IsZero() {
        t.Error("Expected non-zero time, got zero time")
    }

    // Calculate offset
    offset := solarTime.Sub(testTime)

    // Verify offset is within reasonable bounds (should be less than 12 hours)
    if offset.Hours() > 12 || offset.Hours() < -12 {
        t.Errorf("Expected offset within ±12 hours, got %f hours", offset.Hours())
    }
}
