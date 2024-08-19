package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./pages/index.html")
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse the index.html")
		return
	}

	tmp, err = tmp.ParseGlob("./templates/*.html")
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse all templates")
		return
	}

	var artists []data.ArtistType

	err = funcs.GetAndParse(data.MainData.Artists, &artists)
	if err != nil {
		ErrorHandler(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	minmaxCreation, err := data.GetMinMaxCreationDate(artists)
	if err != nil {
		ErrorHandler(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	if len(r.URL.Query()) != 0 {
		artists = data.FilterArtists(artists, r.URL.Query())
	}
	minCreationValue := strconv.Itoa(minmaxCreation["min"])
	maxCreationValue := strconv.Itoa(minmaxCreation["max"])
	if len(r.URL.Query()["min-creation"]) != 0 {
		minCreationValue = r.URL.Query()["min-creation"][0]
	}
	if len(r.URL.Query()["max-creation"]) != 0 {
		maxCreationValue = r.URL.Query()["max-creation"][0]
	}

	homeData := struct {
		Artists []data.ArtistType
		Filter  data.FilterType
	}{
		Artists: artists,
		Filter: data.FilterType{
			Creation: struct {
				MinCreationDate int
				MaxCreationDate int
				MinValue        string
				MaxValue        string
			}{
				MinCreationDate: minmaxCreation["min"],
				MaxCreationDate: minmaxCreation["max"],
				MinValue:        minCreationValue,
				MaxValue:        maxCreationValue,
			},
		},
	}

	err = tmp.Execute(w, homeData)
	if err != nil {
		fmt.Println("When we excute the html", err)
		return
	}
}
