package serverapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	Next     = "next"
	Previous = "previous"
)

// RespLocationInfo contains information about user's current location.
func (c *Client) LocationInfo() (RespLocationInfo, error) {
	if c.token == "" {
		return RespLocationInfo{}, fmt.Errorf("log in to explore the world")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/location", nil)
	if err != nil {
		return RespLocationInfo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = bodyToError(resp.Body)
		return RespLocationInfo{}, err
	}

	var response RespLocationInfo
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return RespLocationInfo{}, err
	}

	return response, nil
}

// Move moves the user to the next or previous location.
func (c *Client) Move(direction string) (RespLocationInfo, error) {
	if c.token == "" {
		return RespLocationInfo{}, fmt.Errorf("log in to explore the world")
	}

	if direction != Next && direction != Previous {
		return RespLocationInfo{}, fmt.Errorf("invalid direction. Use 'next' or 'previous'")
	}

	req, err := http.NewRequest("PUT", c.baseURL+"/location/"+direction, nil)
	if err != nil {
		return RespLocationInfo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = bodyToError(resp.Body)
		return RespLocationInfo{}, err
	}

	return c.LocationInfo()
}

func (c *Client) Search() (Pokemon, error) {
	if c.token == "" {
		return Pokemon{}, fmt.Errorf("log in to explore the world")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/location/search", nil)
	if err != nil {
		return Pokemon{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = bodyToError(resp.Body)
		return Pokemon{}, err
	}

	var response Pokemon
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Pokemon{}, err
	}

	return response, nil
}
