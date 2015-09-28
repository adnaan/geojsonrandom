# geojsonrandom
Generates random geojson points. [godoc](https://godoc.org/github.com/adnaan/geojsonrandom).

go get github.com/adnaan/geojsonrandom

...
Generate points in a circle

points := geojsonrandom.RandomPointsInCircle(10, geojsonrandom.BANGALORE, 5000)

...
Generate points in a bounding box

bbox := []float64{-10.0, -10.0, 10.0, 10.0}
points := RandomPointsInBox(10, bbox))
