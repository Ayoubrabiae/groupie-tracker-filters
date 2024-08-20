package data

import (
	"fmt"

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
