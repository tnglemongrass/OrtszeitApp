package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

type IPInfo struct {
	IP                 string  `json:"ip"`
	Network            string  `json:"network"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeISO3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTLD         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEU               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UTCOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	ASN                string  `json:"asn"`
	Org                string  `json:"org"`
}

var karlsruheIPInfo = &IPInfo{
	City:      "Karlsruhe",
	Region:    "Baden-Württemberg",
	Country:   "DE",
	Latitude:  49.009375,
	Longitude: 8.40425,
	Timezone:  "Europe/Berlin",
}

func getIPInfo(debug bool) (*IPInfo, error) {
	resp, err := http.Get("https://ipapi.co/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if debug {
		fmt.Println(string(body)) // Print the response body
	}
	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if debug {
		fmt.Printf("IPInfo: %+v\n", ipInfo) // Debug print IPInfo struct
	}
	if err != nil {
		var errorResponse map[string]interface{}
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return nil, err
		}
		if errorResponse["error"] == true {
			return nil, fmt.Errorf("API error: %s", errorResponse["message"])
		}
		return nil, err
	}

	return &ipInfo, nil
}

// Declination calculates the solar declination angle in degrees for a given day of the year.
func Declination(dayOfYear int) float64 {
	L := (2 * math.Pi / 365.25) * float64(dayOfYear-81)
	declination := 23.45 * math.Sin(L)
	return declination
}

func equationOfTime(dayOfYear int) float64 {
	L := (2 * math.Pi / 365.25) * float64(dayOfYear-81)
	return (229.18*math.Sin(2*L) - 7.53*math.Sin(4*L) + 2.89*math.Sin(6*L)) / 60.0
}

func meanSolarTime(standardTime time.Time, equationOfTime float64) time.Time {
	return standardTime.Add(time.Duration(equationOfTime * float64(time.Hour)))
}

func apparentSolarTime(meanSolarTime time.Time, declination float64) time.Time {
	return meanSolarTime.Add(time.Duration(4.0 * math.Sin(declination) * float64(time.Minute)))
}

func main() {
	karlsruhe := flag.Bool("karlsruhe", false, "use Karlsruhe, Germany")
	debug := flag.Bool("debug", false, "print JSON response")
	flag.Parse()

	var ipInfo *IPInfo
	var err error

if *karlsruhe {
	ipInfo = karlsruheIPInfo
} else {
	ipInfo, err = getIPInfo(*debug)
	if err != nil {
		fmt.Println(err)
		return
	}
}

	// Calculate standard time based on timezone
	utcTime := time.Now().UTC()
	location, err := time.LoadLocation(ipInfo.Timezone)
	if err != nil {
		fmt.Println("Error loading timezone:", err)
		return
	}
	standardTime := utcTime.In(location)

	dayOfYear := standardTime.YearDay()

	eot := equationOfTime(dayOfYear)
	meanSolarTime := meanSolarTime(standardTime, eot)

	declination := Declination(dayOfYear)
	apparentSolarTime := apparentSolarTime(meanSolarTime, declination)

	fmt.Printf("Location:           %s, %s, %s\n", ipInfo.City, ipInfo.Region, ipInfo.Country)
	fmt.Printf("Coordinates:        %.4f°N, %.4f°E\n", ipInfo.Latitude, ipInfo.Longitude)
	fmt.Printf("Timezone:           %s (UTC%s)\n", ipInfo.Timezone, ipInfo.UTCOffset)

	fmt.Printf("UTC:                %s\n", time.Now().UTC().Format("15:04:05 / 2006-01-02"))
	fmt.Printf("Local:              %s\n", standardTime.Format("15:04:05"))

	fmt.Printf("\nCorrections:\n")
	fmt.Printf("Equation of time:   %.3f minutes\n", eot)

	fmt.Printf("\nTime Calculations:\n")
	fmt.Printf("Mean solar time:    %s\n", meanSolarTime.Format("15:04:05"))
	fmt.Printf("Apparent solar time:%s\n", apparentSolarTime.Format("15:04:05"))

}
