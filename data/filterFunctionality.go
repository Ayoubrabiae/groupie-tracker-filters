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

// Get min and max values for ranges inputs
func GetMinMaxValues(minmax map[string]int, min, max []string) (string, string) {
	minValue := strconv.Itoa(minmax["min"])
	maxValue := strconv.Itoa(minmax["max"])

	if len(min) != 0 && len(max) != 0 {
		minInt, _ := strconv.Atoi(min[0])
		maxInt, _ := strconv.Atoi(max[0])

		if minInt > minmax["max"] || minInt < minmax["min"] {
			return minValue, maxValue
		}
		if maxInt > minmax["max"] || maxInt < minmax["min"] {
			return minValue, maxValue
		}

		minValue = min[0]
		maxValue = max[0]
	}

	return minValue, maxValue
}

// Get Members sizes from artists
func GetMembersSizes(artists []ArtistType) map[int]bool {
	res := map[int]bool{}

	for _, artist := range artists {
		res[len(artist.Members)] = true
	}

	return res
}

// Get Members Sizes That we check
func GetCheckedMembers(members []string) map[int]bool {
	res := map[int]bool{}

	for _, m := range members {
		mInt, err := strconv.Atoi(m)
		if err != nil {
			return map[int]bool{}
		}

		res[mInt] = true
	}

	return res
}

// Filter Artists based on members size
func MembersFilter(artists []ArtistType, membersSizes []string) ([]ArtistType, error) {
	res := []ArtistType{}

	for _, artist := range artists {
		for _, size := range membersSizes {

			intSize, err := strconv.Atoi(size)
			if err != nil {
				return []ArtistType{}, err
			}

			if len(artist.Members) == intSize {
				res = append(res, artist)
			}

		}
	}

	return res, nil
}

// Help us to filter Artists based on the creation and the first album date
func CreationFilter(artists []ArtistType, minCreaionDate, maxCreationDate string) ([]ArtistType, error) {
	if minCreaionDate == "" || maxCreationDate == "" {
		return artists, nil
	}

	var res []ArtistType

	min, err := strconv.Atoi(minCreaionDate)
	if err != nil {
		return []ArtistType{}, err
	}

	max, err := strconv.Atoi(maxCreationDate)
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

func FirstAlbumFilter(artists []ArtistType, minFirstAlbum, maxFirstAlbum string) ([]ArtistType, error) {
	if minFirstAlbum == "" || maxFirstAlbum == "" {
		return artists, nil
	}

	var res []ArtistType

	min, err := strconv.Atoi(minFirstAlbum)
	if err != nil {
		return []ArtistType{}, err
	}

	max, err := strconv.Atoi(maxFirstAlbum)
	if err != nil {
		return []ArtistType{}, err
	}

	for _, item := range artists {
		firstAlbumDate, err := funcs.DateToInt(item.FirstAlbum)
		if err != nil {
			return []ArtistType{}, err
		}

		if firstAlbumDate >= min && firstAlbumDate <= max {
			res = append(res, item)
		}
	}

	return res, nil
}

// Get Selected Location
func GetSelectedLocation(location []string) string {
	if len(location) == 0 {
		return ""
	}

	return location[0]
}

// Filter Artists Based on locations
func LocationFilter(artists []ArtistType, location string) ([]ArtistType, error) {
	if location == "" {
		return artists, nil
	}

	var locationsHolder struct {
		Index []struct {
			LocationsType
		} `json:"index"`
	}

	err := funcs.GetAndParse(MainData.Locations, &locationsHolder)
	if err != nil {
		return []ArtistType{}, nil
	}

	theIds := map[int]bool{}

	for _, locStruct := range locationsHolder.Index {
		for _, loc := range locStruct.Locations {
			if loc == location {
				theIds[locStruct.Id] = true
			}
		}
	}

	res := []ArtistType{}

	for _, artist := range artists {
		if theIds[artist.Id] {
			res = append(res, artist)
		}
	}

	return res, nil
}

// Filter Artists Used Functions Above
func FilterArtists(artists []ArtistType, p map[string][]string) []ArtistType {
	var res []ArtistType = artists
	var err error

	if len(p["location"]) != 0 {
		res, err = LocationFilter(res, p["location"][0])
		if err != nil {
			return []ArtistType{}
		}
	}

	if len(p["members"]) != 0 {
		res, err = MembersFilter(res, p["members"])
		if err != nil {
			return []ArtistType{}
		}
	}

	if len(p["min-creation"]) != 0 && len(p["max-creation"]) != 0 {
		res, err = CreationFilter(res, p["min-creation"][0], p["max-creation"][0])
		if err != nil {
			return []ArtistType{}
		}
	}

	if len(p["min-first-album"]) != 0 && len(p["max-first-album"]) != 0 {
		res, err = FirstAlbumFilter(res, p["min-first-album"][0], p["max-first-album"][0])
		if err != nil {
			return []ArtistType{}
		}
	}

	return res
}

// Get All Filter Params Like range values and checkbox that we checked
func GetFilterParams(artists []ArtistType, p map[string][]string) (FilterType, error) {
	minmaxCreation, err := GetMinMaxCreationDate(artists)
	if err != nil {
		return FilterType{}, err
	}

	minCreationValue, maxCreationValue := GetMinMaxValues(minmaxCreation, p["min-creation"], p["max-creation"])

	//////////////////////////////////////////////////////////

	minmaxFirstAlbum, err := GetMinMaxFirstAlbum(artists)
	if err != nil {
		return FilterType{}, err
	}

	minFirstAlbumValue, maxFirstAlbumValue := GetMinMaxValues(minmaxFirstAlbum, p["min-first-album"], p["max-first-album"])

	//////////////////////////////////////////////////////////

	membersSizes := GetMembersSizes(artists)
	checkedMembers := GetCheckedMembers(p["members"])

	//////////////////////////////////////////////////////////

	locations, err := LoadLocations()
	selectedLocation := GetSelectedLocation(p["location"])
	if err != nil {
		return FilterType{}, err
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
		MembersFilter: MembersFilterType{
			MembersSizes:   membersSizes,
			MembersChecked: checkedMembers,
		},
		LocationsFilter: LocationsFilterType{
			Locations:       locations,
			LocationChecked: selectedLocation,
		},
	}, nil
}
