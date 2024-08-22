package data

type MainType struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type ArtistType struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type LocationsType struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesType struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationsType struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistInfo struct {
	Artist    ArtistType
	Locations LocationsType
	Dates     DatesType
	Relations RelationsType
}

type CreationFilterType struct {
	Min      int
	Max      int
	MinValue string
	MaxValue string
}

type FirstAlbumFilterType struct {
	Min      int
	Max      int
	MinValue string
	MaxValue string
}

type MembersFilterType struct {
	MembersSizes   map[int]bool
	MembersChecked map[int]bool
}

type LocationsFilterType struct {
	Locations       map[string]bool
	LocationChecked string
}

type FilterType struct {
	CreationFilter   CreationFilterType
	FirstAlbumFilter FirstAlbumFilterType
	MembersFilter    MembersFilterType
	LocationsFilter  LocationsFilterType
}

var (
	MainData       MainType
	CoordinatesApi string = "https://photon.komoot.io/api/?q="
	MapLink               = "https://www.google.com/maps/search/?api=1&query=%s"
)
