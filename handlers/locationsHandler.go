package handlers

import (
	"net/http"
	"regexp"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`/locations/([\w|-]+)$`)

	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	place := matches[1]
	allowdedPlaces, err := data.LoadLocations()
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !allowdedPlaces[place] {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	cords, err := funcs.GetCoordinates(data.CoordinatesApi + funcs.FixStringForApi(place) + "&limit=1")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	link, err := funcs.MakeLocationLink(data.MapLink, cords[0], cords[1])
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}
