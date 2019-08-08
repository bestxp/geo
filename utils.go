package geo

import "math"

const (
	// EarthRadius Earth's radius (at the Equator) of 6378137 meters.
	EarthRadius = 6378137
)

func toDegrees(radians float64) float64 {
	return (radians * 180) / math.Pi
}

func toRadians(angleDegrees float64) float64 {
	return (angleDegrees * math.Pi) / 180.0
}
