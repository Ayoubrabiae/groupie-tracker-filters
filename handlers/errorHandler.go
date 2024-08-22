package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, errorMsg string, code int) {
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("./pages/error.html")
	if err != nil {
		http.Error(w, "internal server error", 500)
		fmt.Println("error when we parse", err)
		return
	}

	errorData := struct {
		Message string
		Code    int
	}{errorMsg, code}

	err = tmpl.Execute(w, errorData)
	if err != nil {
		fmt.Println("Error When Excute", err)
		return
	}
}
