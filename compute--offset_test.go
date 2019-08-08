package geo

import "testing"

func TestComputeOffset(t *testing.T) {

	p := Point{55.623151, 8.48215}
	testPoint := Point{55.66806676420599, 8.48215}

	newP := ComputeOffset(p, 5000, 0)

	if testPoint.Lat != newP.Lat {
		t.Errorf("Lat: %v not equal %v", testPoint.Lat, newP.Lat)
	}
	if testPoint.Lng != newP.Lng {
		t.Errorf("Lng: %v not equal %v", testPoint.Lng, newP.Lng)
	}
}
