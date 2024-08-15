package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
	"groupie-tracker/handlers"
)

func TestProfilehandler(t *testing.T) {
	err := funcs.GetAndParse("https://groupietrackers.herokuapp.com/api", &data.MainData)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/artists/1", nil)
	if err != nil {
		t.Error(err)
	}
	handlers.ProfileHandler(resp, req)
	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("expected status 200 got %d", resp.Result().StatusCode)
	}
	defer resp.Result().Body.Close()
}

func TestHomehandler(t *testing.T) {
	data.MainData.Artists = "https://groupietrackers.herokuapp.com/api/artists"
	server := httptest.NewServer(http.HandlerFunc(handlers.HomeHandler))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200 got %d", resp.StatusCode)
	}
}
