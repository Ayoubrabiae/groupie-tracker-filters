package funcs

import (
	"errors"
	"fmt"
	"strings"
)

func FixStringForApi(s string) string {
	s = strings.ReplaceAll(s, "_", "-")
	s = strings.ReplaceAll(s, "-usa", "")

	return s
}

func GetCoordinates(url string) ([]float64, error) {
	cords := struct {
		Features []struct {
			Geometry struct {
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"features"`
	}{}

	err := GetAndParse(url, &cords)
	if err != nil {
		return []float64{}, err
	}

	if len(cords.Features) < 1 {
		return []float64{}, errors.New("invalid url")
	}

	return cords.Features[0].Geometry.Coordinates, nil
}

func MakeLocationLink(mapLink string, lon, lat float64) (string, error) {
	cords := fmt.Sprintf("%.6f,%.6f", lat, lon)

	link := fmt.Sprintf(mapLink, cords)

	return link, nil
}
