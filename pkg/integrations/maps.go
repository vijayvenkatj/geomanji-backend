package integrations

import "fmt"

func GoogleMapsLink(lat, lng float64) string {
	return fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", lat, lng)
}

// GoogleEarthLink builds a Google Earth Web URL pointed at the given
// coordinates with a tilted camera (3D perspective) and a tight zoom;
// Earth animates the fly-in itself when the link is opened. No API key
// required. Camera params: altitude,distance,fov,heading,tilt,roll.
func GoogleEarthLink(lat, lng float64) string {
	return fmt.Sprintf("https://earth.google.com/web/@%f,%f,0a,300d,35y,45h,65t,0r", lat, lng)
}
