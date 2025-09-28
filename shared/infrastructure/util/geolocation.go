package util

import (
	"strconv"
	"strings"
	"team-maker-api/shared/model/vo"
)

func GetLatlngString(latitude string, longitude string) (vo.Geolocation, error) {
	var LatLng vo.Geolocation
	// LATITUDE
	Latitude, err := strconv.ParseFloat(strings.TrimSpace(latitude), 64)
	if err != nil {
		return LatLng, err
	}
	// LONGITUDE
	Longitude, err := strconv.ParseFloat(strings.TrimSpace(longitude), 64)
	if err != nil {
		return LatLng, err
	}
	LatLng[0] = Latitude
	LatLng[1] = Longitude

	return LatLng, nil
}
