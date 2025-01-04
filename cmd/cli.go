package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
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

func main() {
	debug := flag.Bool("debug", false, "print JSON response")
	flag.Parse()

	ipInfo, err := getIPInfo(*debug)
	if err != nil {
		fmt.Println(err)
		return
	}

fmt.Printf("Current Time (UTC): %s\n", time.Now().UTC().Format("2006-01-02 15:04:05"))
fmt.Printf("Location:           %s, %s, %s\n", ipInfo.City, ipInfo.Region, ipInfo.Country)
fmt.Printf("Coordinates:        %.4f°N, %.4f°E\n", ipInfo.Latitude, ipInfo.Longitude)
fmt.Printf("Timezone:           %s (UTC%s)\n", ipInfo.Timezone, ipInfo.UTCOffset)
}
