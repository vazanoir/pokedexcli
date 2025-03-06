package client

import (
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) GetLocations(url string) (LocationPage, error) {
	if body, found := c.cache.Get(url); found {
		page := LocationPage{}
		if err := json.Unmarshal(body, &page); err != nil {
			return LocationPage{}, err
		}
		return page, nil
	}

	res, err := c.httpClient.Get(url)
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

	c.cache.Add(url, body)

	return page, nil
}

type ExploreResult struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ExploreLocations(url string) (ExploreResult, error) {
	if body, found := c.cache.Get(url); found {
		page := ExploreResult{}
		if err := json.Unmarshal(body, &page); err != nil {
			return ExploreResult{}, err
		}
		return page, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return ExploreResult{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return ExploreResult{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	if err != nil {
		return ExploreResult{}, err
	}

	result := ExploreResult{}
	if err := json.Unmarshal(body, &result); err != nil {
		return ExploreResult{}, err
	}

	c.cache.Add(url, body)

	return result, nil
}

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
	Weight         int    `json:"weight"`
	Height         int    `json:"height"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) CatchPokemon(url string) (Pokemon, error) {
	if body, found := c.cache.Get(url); found {
		pokemon := Pokemon{}
		if err := json.Unmarshal(body, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return pokemon, nil
}
