package data

import (
	"errors"
	"strconv"

	"groupie-tracker/funcs"
)

// Get min and max creation dates from artists
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

// Get min and max first album dates from artists
func GetMinMaxFirstAlbum(artists []ArtistType) (map[string]int, error) {
	res := map[string]int{}
	if len(artists) < 1 {
		return res, errors.New("artists slice is empty")
	}

	var err error
	res["min"], err = funcs.DateToInt(artists[0].FirstAlbum)
	if err != nil {
		return map[string]int{}, err
	}

	for _, artist := range artists {
		year, err := funcs.DateToInt(artist.FirstAlbum)
		if err != nil {
			return map[string]int{}, err
		}

		if res["min"] > year {
			res["min"] = year
		} else if res["max"] < year {
			res["max"] = year
		}
	}

	return res, nil
}

// Filter Artists based on the creation date
func CreationFilter(artists []ArtistType, minCreaionDate, maxCreationDate []string) ([]ArtistType, error) {
	res := []ArtistType{}

	if len(minCreaionDate) == 0 || len(maxCreationDate) == 0 {
		return []ArtistType{}, nil
	}

	min, err := strconv.Atoi(minCreaionDate[0])
	if err != nil {
		return []ArtistType{}, err
	}

	max, err := strconv.Atoi(maxCreationDate[0])
	if err != nil {
		return []ArtistType{}, err
	}

	for _, item := range artists {
		if item.CreationDate >= min && item.CreationDate <= max {
			res = append(res, item)
		}
	}

	return res, nil
}

// Filter Artists Used Functions Above
func FilterArtists(artists []ArtistType, p map[string][]string) ([]ArtistType, error) {
	res, err := CreationFilter(artists, p["min-creation"], p["max-creation"])
	if err != nil {
		return []ArtistType{}, err
	}

	return res, nil
}

// Get All Filter Params Like range values and checkbox that we checked
func GetFilterParams(artists []ArtistType, p map[string][]string) (FilterType, error) {
	minmaxCreation, err := GetMinMaxCreationDate(artists)
	if err != nil {
		return FilterType{}, err
	}

	minCreationValue := strconv.Itoa(minmaxCreation["min"])
	maxCreationValue := strconv.Itoa(minmaxCreation["max"])
	if len(p["min-creation"]) != 0 {
		minCreationValue = p["min-creation"][0]
	}
	if len(p["max-creation"]) != 0 {
		maxCreationValue = p["max-creation"][0]
	}

	//////////////////////////////////////////////////////////

	minmaxFirstAlbum, err := GetMinMaxFirstAlbum(artists)
	if err != nil {
		return FilterType{}, err
	}

	minFirstAlbumValue := strconv.Itoa(minmaxFirstAlbum["min"])
	maxFirstAlbumValue := strconv.Itoa(minmaxFirstAlbum["max"])
	if len(p["min-first-album"]) != 0 {
		minCreationValue = p["min-first-album"][0]
	}
	if len(p["max-first-album"]) != 0 {
		maxCreationValue = p["max-first-album"][0]
	}
	
	//////////////////////////////////////////////////////////

	return FilterType{
		CreationFilter: CreationFilterType{
			Min:      minmaxCreation["min"],
			Max:      minmaxCreation["max"],
			MinValue: minCreationValue,
			MaxValue: maxCreationValue,
		},
		FirstAlbumFilter: FirstAlbumFilterType{
			Min:      minmaxFirstAlbum["min"],
			Max:      minmaxFirstAlbum["max"],
			MinValue: minFirstAlbumValue,
			MaxValue: maxFirstAlbumValue,
		},
	}, nil
}
