package geojsonrandom

import (
	"math"
	"math/rand"
)

// Geometry types
const (
	point string = "point"
)

// Geometry geoson geometry
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Where is Sodom and Gommorrah?
var (
	NEWYORK   = []float64{40.750422, -73.996328}
	BANGALORE = []float64{12.953997, 77.630939}
)

// RandomPointsInCircle generates random points in a circle
func RandomPointsInCircle(count int, center []float64, radiusInMeters float64) []Geometry {
	var geometryArr []Geometry

	radiusInDegrees := radiusInMeters / 111300
	for index := 0; index < count; index++ {

		u := rand.Float64()
		v := rand.Float64()

		w := radiusInDegrees * math.Sqrt(u)
		t := 2 * math.Pi * v
		x := float64(int((w*math.Cos(t)+center[0])*1e6)) / 1e6
		y := float64(int((w*math.Sin(t)+center[1])*1e6)) / 1e6

		geometryArr = append(geometryArr, Geometry{Type: point, Coordinates: []float64{x, y}})
	}
	return geometryArr
}

// RandomPointsInBox generates random points in a bounding box
func RandomPointsInBox(count int, bbox []float64) []Geometry {
	var geometryArr []Geometry
	for index := 0; index < count; index++ {
		u := (rand.Float64() * (bbox[2] - bbox[0])) + bbox[0]
		v := (rand.Float64() * (bbox[2] - bbox[0])) + bbox[0]
		x := float64(int(u*1e6)) / 1e6
		y := float64(int(v*1e6)) / 1e6
		geometryArr = append(geometryArr, Geometry{Type: point, Coordinates: []float64{x, y}})
	}
	return geometryArr
}
