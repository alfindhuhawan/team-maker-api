package vo

import (
	"math"
	"strconv"
)

type Geolocation [2]float64

func (g Geolocation) GetLatitude() float64 {
	return g[0]
}

func (g Geolocation) GetLongitude() float64 {
	return g[1]
}

func (g Geolocation) GetLatitudeString() string {
	return strconv.FormatFloat(g.GetLatitude(), 'f', 6, 64)
}

func (g Geolocation) GetLongitudeSting() string {
	return strconv.FormatFloat(g.GetLongitude(), 'f', 6, 64)
}

func (g Geolocation) GetDistance(a Geolocation) float64 {
	return math.Sqrt(g.GetLongitude()*g.GetLongitude() + g.GetLatitude()*g.GetLatitude())
}
