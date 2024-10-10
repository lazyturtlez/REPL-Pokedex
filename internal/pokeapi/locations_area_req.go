package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit !")
		locationData := LocationAreasResponse{}
		if err := json.Unmarshal(dat, &locationData); err != nil {
			return LocationAreasResponse{}, err
		}
		return locationData, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationData := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationData)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	fmt.Println("cache miss!")
	c.cache.Add(fullURL, dat)
	return locationData, nil
}

func (c *Client) ListLocationData(pokeLocation string) (LocationAreaData, error){
	if pokeLocation == "" {
		return LocationAreaData{}, fmt.Errorf("please provide a location to search")
	}
	endpoint := "/location-area/"
	fullURL := baseURL+endpoint+pokeLocation

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit !")
		locationAreaData := LocationAreaData{}
		if err := json.Unmarshal(dat, &locationAreaData); err != nil {
			return LocationAreaData{}, err
		}
		c.cache.Add(fullURL, dat)
		return locationAreaData, nil
	}


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaData{}, err
	}
	defer res.Body.Close()

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaData{}, err
	}
	locationAreaData := LocationAreaData{}

	err = json.Unmarshal(dat, &locationAreaData)
	if err != nil {
		return LocationAreaData{}, err
	}
	fmt.Println("cache miss !")
	c.cache.Add(fullURL, dat)
	return locationAreaData, nil
}