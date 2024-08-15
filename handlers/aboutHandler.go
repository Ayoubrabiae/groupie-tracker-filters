package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./pages/about.html")
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

	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Println("When we excute", err)
		return
	}
}
