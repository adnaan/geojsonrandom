package geosonrandom

import (
	"fmt"
	"testing"
)

func TestPoints(t *testing.T) {

	fmt.Println("points in a cirlce")
	fmt.Println(RandomPointsInCircle(10, NEWYORK, 5000))
	bbox := []float64{-10.0, -10.0, 10.0, 10.0}
	fmt.Println("points in a box")
	fmt.Println(RandomPointsInBox(10, bbox))

}
