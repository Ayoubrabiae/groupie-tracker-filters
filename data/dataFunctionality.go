package data

import (
	"errors"
	"fmt"
	"strconv"

	"groupie-tracker/funcs"
)

func LoadArtistData(id string) (ArtistInfo, error) {
	var (
		artist    ArtistType
		dates     DatesType
		locations LocationsType
		relations RelationsType
	)

	var err error

	err = funcs.GetAndParse(MainData.Artists+"/"+id, &artist)
	if err != nil {
		return ArtistInfo{}, err
	}

	err = funcs.GetAndParse(MainData.Dates+"/"+id, &dates)
	if err != nil {
		return ArtistInfo{}, err
	}

	err = funcs.GetAndParse(MainData.Locations+"/"+id, &locations)
	if err != nil {
		fmt.Println(err)
		return ArtistInfo{}, err
	}

	err = funcs.GetAndParse(MainData.Relations+"/"+id, &relations)
	if err != nil {
		return ArtistInfo{}, err
	}

	return ArtistInfo{
		Artist:    artist,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}, nil
}

func IdCheck(id string) (bool, error) {
	idHolder := struct {
		Id int `json:"id"`
	}{}

	err := funcs.GetAndParse(MainData.Artists+"/"+id, &idHolder)
	if err != nil {
		return false, err
	}

	return idHolder.Id != 0, nil
}

func LoadLocations() (map[string]bool, error) {
	var locationsHolder struct {
		Index []struct {
			Locations []string `json:"locations"`
		} `json:"index"`
	}
	err := funcs.GetAndParse(MainData.Locations, &locationsHolder)
	if err != nil {
		return map[string]bool{}, err
	}

	res := map[string]bool{}
	for _, item := range locationsHolder.Index {
		for _, element := range item.Locations {
			res[element] = true
		}
	}

	return res, nil
}

func GetMinMaxCreationDate(artists []ArtistType) (map[string]int, error) {
	res := map[string]int{}
	if len(artists) < 1 {
		return res, errors.New("artists slice is empty")
	}

	res["min"] = artists[0].CreationDate

	for _, artist := range artists {
		if artist.CreationDate < res["min"] {
			res["min"] = artist.CreationDate
		} else if artist.CreationDate > res["max"] {
			res["max"] = artist.CreationDate
		}
	}

	return res, nil
}

func FilterArtists(artists []ArtistType, p map[string][]string) []ArtistType {
	res := []ArtistType{}

	min, _ := strconv.Atoi(p["min-creation"][0])
	max, _ := strconv.Atoi(p["max-creation"][0])

	for _, item := range artists {
		if item.CreationDate >= min && item.CreationDate <= max {
			res = append(res, item)
		}
	}

	return res
}
