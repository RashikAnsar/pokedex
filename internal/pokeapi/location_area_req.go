package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache Hit")
		fmt.Println(fullURL)
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil
	}

	fmt.Println("Cache MISS")
	fmt.Println(fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, nil
	}

	c.cache.Add(fullURL, data)
	return locationAreaResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache Hit")
		fmt.Println(fullURL)
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreaResp, nil
	}

	fmt.Println("Cache MISS")
	fmt.Println(fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResp := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationArea{}, nil
	}

	c.cache.Add(fullURL, data)
	return locationAreaResp, nil
}
