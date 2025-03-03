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
