package coordinates

import (
	iss "issposlib"
	"location"
	"math"
)

const (
	EquatorialRadius = 6378
	PolarRadius      = 6356
	DegreesToRadians = 0.0174532925
	RadiansToDegrees = 57.295779513
)

func Square(v float64) float64 {
	return v * v
}

func EarthRadiusAt(coords *location.GPSCoords) float64 {
	latCos := math.Cos(coords.Latitude * DegreesToRadians)
	lonSin := math.Sin(coords.Longitude * DegreesToRadians)
	return math.Sqrt((Square(Square(EquatorialRadius)*latCos) + Square(Square(PolarRadius)*lonSin)) / (Square(EquatorialRadius*latCos) + Square(PolarRadius*lonSin)))
}

func ToCartesian(coords *location.GPSCoords, radiusExtent float64) (float64, float64, float64) {
	r := EarthRadiusAt(coords) + radiusExtent
	x := r * math.Sin(coords.Latitude*DegreesToRadians) * math.Cos(coords.Longitude*DegreesToRadians)
	y := r * math.Sin(coords.Latitude*DegreesToRadians) * math.Sin(coords.Longitude*DegreesToRadians)
	z := r * math.Cos(coords.Longitude*DegreesToRadians)

	return x, y, z
}

func ToRelative(interest *iss.ISSPosition, observer *location.GPSCoords) (float64, float64) {
	xo, yo, zo := ToCartesian(observer, 0)
	xi, yi, zi := ToCartesian(&interest.GPSCoords, interest.Altitude)

	xr := xi - xo
	yr := yi - yo
	zr := zi - zo

	// Calculate the azimuthal angle
	phi := math.Atan2(yr, xr) * RadiansToDegrees

	// Calculate the inclination angle
	theta := math.Acos(zr/(math.Sqrt(Square(xr)+Square(yr)+Square(zr)))) * RadiansToDegrees

	return phi, theta
}
