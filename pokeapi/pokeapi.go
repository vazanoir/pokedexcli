package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/vazanoir/pokedexcli/cache"
	"io"
	"net/http"
	"time"
)

type LocationPage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  [20]struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

var c cache.Cache = cache.NewCache(7 * time.Second)

func GetLocations(url string) (LocationPage, error) {
	if body, found := c.Get(url); found {
        page := LocationPage{}
        if err := json.Unmarshal(body, &page); err != nil {
            return LocationPage{}, err
        }
        return page, nil
    }

    fmt.Printf("INFO: cache not used for '%v'\n", url)
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

	c.Add(url, body)

	return page, nil
}
