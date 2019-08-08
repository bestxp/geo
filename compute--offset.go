package geo

import "math"

// ComputeOffset Returns the LatLng resulting from moving a distance from an origin in the specified heading (expressed in degrees clockwise from north)
// from start point
// distance in meters
// heading degrees clockwise from north
func ComputeOffset(
	from Point,
	distance float64,
	heading float64,
) Point {
	distance = distance / EarthRadius
	heading = toRadians(heading)
	fromLat := toRadians(from.Lat)
	cosDistance := math.Cos(distance)
	sinDistance := math.Sin(distance)
	sinFromLat := math.Sin(fromLat)
	cosFromLat := math.Cos(fromLat)
	sc := cosDistance*sinFromLat + sinDistance*cosFromLat*math.Cos(heading)
	return Point{
		Lat: toDegrees(math.Asin(sc)),
		Lng: toDegrees(toRadians(from.Lng) + math.Atan2(sinDistance*cosFromLat*math.Sin(heading), cosDistance-sinFromLat*sc)),
	}
}

// Point ...
type Point struct {
	Lat float64
	Lng float64
}
