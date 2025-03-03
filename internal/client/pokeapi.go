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
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
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

	page := ExploreResult{}
	if err := json.Unmarshal(body, &page); err != nil {
		return ExploreResult{}, err
	}

	c.cache.Add(url, body)

	return page, nil
}
