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
	if len(min) != 0 {
		minValue = min[0]
	}
	if len(max) != 0 {
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
func GetCheckedMembers(members []string) (map[int]bool, error) {
	res := map[int]bool{}

	for _, m := range members {
		mInt, err := strconv.Atoi(m)
		if err != nil {
			return map[int]bool{}, err
		}

		res[mInt] = true
	}

	return res, nil
}

// Filter Artists based on members size
func MembersFilter(artists []ArtistType, membersSizes []string) ([]ArtistType, error) {
	res := []ArtistType{}

	if len(membersSizes) == 0 {
		return artists, nil
	}

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
func RangeFilter(artists []ArtistType, minCreaionDate, maxCreationDate []string) ([]ArtistType, error) {
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
func FilterArtists(artists []ArtistType, p map[string][]string) []ArtistType {
	res, err := RangeFilter(artists, p["min-creation"], p["max-creation"])
	if err != nil {
		return []ArtistType{}
	}

	res, err = RangeFilter(res, p["min-first-album"], p["max-first-album"])
	if err != nil {
		return []ArtistType{}
	}

	res, err = MembersFilter(res, p["members"])
	if err != nil {
		return []ArtistType{}
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
	checkedMembers, err := GetCheckedMembers(p["members"])
	if err != nil {
		return FilterType{}, err
	}

	//////////////////////////////////////////////////////////

	locations, err := LoadLocations()
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
			Locations: locations,
		},
	}, nil
}
