package main

import (
	"fmt"
	"math"
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type IPInfo struct {
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}

var karlsruheIPInfo = &IPInfo{
	City:      "Karlsruhe",
	Region:    "Baden-Württemberg",
	Country:   "DE",
	Latitude:  49.009375,
	Longitude: 8.40425,
	Timezone:  "Europe/Berlin",
}

func equationOfTime(dayOfYear int, year int) float64 {
	D := 6.24004077 + 0.01720197*(365.25*float64(year-2000)+float64(dayOfYear))
	return -7.659*math.Sin(D) + 9.863*math.Sin(2*D+3.5932)
}

func meanSolarTime(utcTime time.Time, longitude float64) time.Time {
	localCorrection := longitude * 4.0
	return utcTime.Add(time.Duration(localCorrection * float64(time.Minute)))
}

func apparentSolarTime(meanSolarTime time.Time, equationOfTime float64) time.Time {
	return meanSolarTime.Add(time.Duration(equationOfTime * float64(time.Minute)))
}

func updateTimeLabels(ipInfo *IPInfo, labels []*widget.Label) {
	utcTime := time.Now().UTC()
	standardTime := utcTime.In(time.FixedZone(ipInfo.Timezone, 0))

	dayOfYear := standardTime.YearDay()

	eot := equationOfTime(dayOfYear, standardTime.Year())
	meanSolarTime := meanSolarTime(utcTime, ipInfo.Longitude)
	apparentSolarTime := apparentSolarTime(meanSolarTime, eot)

	labels[0].SetText(fmt.Sprintf("Location:           %s, %s, %s", ipInfo.City, ipInfo.Region, ipInfo.Country))
	labels[1].SetText(fmt.Sprintf("Coordinates:        %.4f°N, %.4f°E", ipInfo.Latitude, ipInfo.Longitude))
	labels[2].SetText(fmt.Sprintf("Timezone:           %s", ipInfo.Timezone))
	labels[3].SetText(fmt.Sprintf("UTC:                %s", utcTime.Format("15:04:05 / 2006-01-02")))
	labels[4].SetText(fmt.Sprintf("Local:              %s", standardTime.Format("15:04:05")))
	labels[5].SetText(fmt.Sprintf("Equation of time:   %.3f minutes", eot))
	labels[6].SetText(fmt.Sprintf("Mean solar time:    %s", meanSolarTime.Format("15:04:05")))
	labels[7].SetText(fmt.Sprintf("Apparent solar time:%s", apparentSolarTime.Format("15:04:05")))
}

func main() {
	a := app.New()
	w := a.NewWindow("Ortszeit")

	ipInfo := karlsruheIPInfo

	labels := []*widget.Label{
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
	}

	updateTimeLabels(ipInfo, labels)

	updateButton := widget.NewButton("Update", func() {
		updateTimeLabels(ipInfo, labels)
	})

	w.SetContent(container.NewVBox(
		labels[0],
		labels[1],
		labels[2],
		labels[3],
		labels[4],
		labels[5],
		labels[6],
		labels[7],
		updateButton,
	))

	w.ShowAndRun()
}
