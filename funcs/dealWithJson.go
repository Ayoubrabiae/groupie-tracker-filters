package funcs

import (
	"encoding/json"
	"net/http"
)

func GetAndParse(url string, v any) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
