package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string
	Url  string
}

type LocationPage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  [20]struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) (LocationPage, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationPage{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationPage{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	if err != nil {
		return LocationPage{}, err
	}

	page := LocationPage{}
	if err := json.Unmarshal(body, &page); err != nil {
		return LocationPage{}, err
	}

	return page, nil
}
