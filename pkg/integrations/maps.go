package integrations

import "fmt"

func GoogleMapsLink(lat, lng float64) string {
	return fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", lat, lng)
}
