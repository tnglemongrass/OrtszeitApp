//go:build android || ios
// +build android ios

package mainold

import (
	"time"
	
	"fyne.io/fyne/v2/driver/mobile"
)

func (o *ortszeit) initMobile() {
	if mobileApp, ok := o.window.Canvas().(mobile.Device); ok {
		// Request location permission
		mobileApp.RequestPermission(mobile.PermissionLocationWhenInUse)
		
		// Start location updates
		go func() {
			for {
				pos, err := mobileApp.GetLocation()
				if err == nil {
					o.updateLocation(pos.Latitude, pos.Longitude)
				}
				// Update every 30 seconds
				time.Sleep(30 * time.Second)
			}
		}()
	}
}
