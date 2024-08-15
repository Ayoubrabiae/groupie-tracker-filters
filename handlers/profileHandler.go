package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"groupie-tracker/data"
)

var IdPath = regexp.MustCompile(`^/artists/(\d+)$`)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	matches := IdPath.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	id := matches[1]

	ok, err := data.IdCheck(id)
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	if !ok {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("./pages/profile.html")
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse the index.html")
		return
	}

	profileData, err := data.LoadArtistData(id)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, profileData)
	if err != nil {
		fmt.Println("When we excute", err)
		return
	}
}
