package geosonrandom

import (
	"math"
	"math/rand"
)

// GeometryType geojson geometry types
type GeometryType string

// Geometry types
const (
	POINT           GeometryType = "point"
	MULTIPOINT      GeometryType = "multipoint"
	LINESTRING      GeometryType = "linestring"
	MULTILINESTRING GeometryType = "multilinestring"
	POLYGON         GeometryType = "polygon"
	MULTIPOLYGON    GeometryType = "multipolygon"
)

// CoordType ...
type CoordType float64

// Coordinate a single coordinate/position in the order latitude,longitude
type Coordinate [2]CoordType

type position []CoordType

type point position

type multiPoint []point

type lineString []position

type linearRing []position

type multiLineString []lineString

type polygon []linearRing

type multiPolygon []polygon

// Bbox geojson bounding box
type Bbox []CoordType

// Properties geoson properties object
type Properties map[string]interface{}

// Geometry geoson geometry
type Geometry struct {
	Type        GeometryType `json:"type"`
	Coordinates interface{}  `json:"coordinates"`
	Bbox        Bbox         `json:"bbox,omitempty"`
}

// GeometryCollection geoson GeometryCollection
type GeometryCollection struct {
	Type       string     `json:"type"`
	Geometries []Geometry `json:"geometries"`
	Bbox       Bbox       `json:"bbox,omitempty"`
}

// Feature geoson feature
type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Bbox       Bbox       `json:"bbox,omitempty"`
	Properties Properties `json:"properties,omitempty"`
}

// FeatureCollection geoson GeometryCollection
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
	Bbox     Bbox      `json:"bbox,omitempty"`
}

// Where is Sodom and Gommorrah?
var (
	NEWYORK = []float64{40.750422, -73.996328}
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

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

		geometryArr = append(geometryArr, Geometry{Type: POINT, Coordinates: []float64{x, y}})
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
		geometryArr = append(geometryArr, Geometry{Type: POINT, Coordinates: []float64{x, y}})
	}
	return geometryArr
}

// // RandomGeometry generate an array of geojson geometries
// func randomGeometry(geometryType GeometryType, bbox Bbox) Geometry {
// 	var coordinates []Coordinate
//
// 	return Geometry{Type: geometryType, Bbox: bbox, Coordinates: coordinates}
// }

// // RandomGeometries generates an array of geojson geometries
// func RandomGeometries(geometryType GeometryType, count int) []Geometry {
// 	var bbox Bbox
// 	return RandomGeometriesBbox(geometryType, count, bbox)
// }
//
// func RandomGeometries(geometryType GeometryType, count int) []Geometry {
// 	var bbox Bbox
// 	return RandomGeometriesBbox(geometryType, count, bbox)
// }
//
// // RandomGeometriesBbox generates an array of geojson geometries with a bounding box
// func RandomGeometriesBbox(geometryType GeometryType, count int, bbox Bbox) []Geometry {
// 	var geometryArr []Geometry
// 	for index := 0; index < count; index++ {
// 		geometryArr = append(geometryArr, randomGeometry(geometryType, bbox))
// 	}
// 	return geometryArr
// }
//
// func randomFeature(geometryType GeometryType, bbox Bbox, properties Properties) Feature {
//
// 	return Feature{Type: "feature", Bbox: bbox, Properties: properties, Geometry: randomGeometry(geometryType, bbox)}
// }
//
// // RandomFeature generate an array of geojson features
// func RandomFeature(geometryType GeometryType, count int, properties Properties) []Feature {
// 	var bbox Bbox
// 	return RandomFeatureBbox(geometryType, count, bbox, properties)
// }
//
// // RandomFeatureBbox generate an array of geojson features with bounding box
// func RandomFeatureBbox(geometryType GeometryType, count int, bbox Bbox, properties Properties) []Feature {
// 	var featureArr []Feature
// 	for index := 0; index < count; index++ {
// 		featureArr = append(featureArr, randomFeature(geometryType, bbox, properties))
// 	}
// 	return featureArr
//
// }
//
// // RandomGeometryCollection generates a geojson GeometryCollection
// func RandomGeometryCollection(geometryTypes []GeometryType) GeometryCollection {
// 	var bbox Bbox
// 	return RandomGeometryCollectionBbox(geometryTypes, bbox)
// }
//
// // RandomGeometryCollectionBbox generates a geojson GeometryCollection with a bounding box
// func RandomGeometryCollectionBbox(geometryTypes []GeometryType, bbox Bbox) GeometryCollection {
// 	var geometries []Geometry
//
// 	return GeometryCollection{Type: "geometrycollection", Geometries: geometries, Bbox: bbox}
// }
//
// // RandomFeatureCollection generates a geojson FeatureCollection
// func RandomFeatureCollection(props map[GeometryType]Properties) FeatureCollection {
// 	var bbox Bbox
// 	return RandomFeatureCollectionBbox(props, bbox)
// }
//
// // RandomFeatureCollectionBbox generates a geojson FeatureCollection with a bounding box
// func RandomFeatureCollectionBbox(props map[GeometryType]Properties, bbox Bbox) FeatureCollection {
// 	var features []Feature
//
// 	return FeatureCollection{Type: "featurecollection", Features: features}
// }
