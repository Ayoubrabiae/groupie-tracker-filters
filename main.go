package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
	"groupie-tracker/handlers"
)

func main() {
	err := funcs.GetAndParse("https://groupietrackers.herokuapp.com/api", &data.MainData)
	if err != nil {
		fmt.Println(err)
		return
	}

	port := "8082"

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static", handlers.StaticHandler(fileServer)))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists/", handlers.ProfileHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/locations/", handlers.LocationsHandler)
	fmt.Println("http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, http.DefaultServeMux))
}
