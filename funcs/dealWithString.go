package funcs

import (
	"errors"
	"strconv"
	"strings"
)

func DateToInt(date string) (int, error) {
	dateSlice := strings.Split(date, "-")

	if len(dateSlice) == 0 {
		return 0, errors.New("please enter a valid date")
	}

	year, err := strconv.Atoi(dateSlice[len(dateSlice)-1])
	if err != nil {
		return 0, err
	}

	return year, nil
}
