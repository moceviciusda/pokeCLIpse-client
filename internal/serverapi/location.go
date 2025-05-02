package serverapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

const (
	Next     = "next"
	Previous = "previous"
)

// RespLocationInfo contains information about user's current location.
func (c *Client) LocationInfo() (RespLocationInfo, error) {
	if c.Token == "" {
		return RespLocationInfo{}, fmt.Errorf("log in to explore the world")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/location", nil)
	if err != nil {
		return RespLocationInfo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

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
	if c.Token == "" {
		return RespLocationInfo{}, fmt.Errorf("log in to explore the world")
	}

	if direction != Next && direction != Previous {
		return RespLocationInfo{}, fmt.Errorf("invalid direction. Use 'next' or 'previous'")
	}

	req, err := http.NewRequest("PUT", c.baseURL+"/location/"+direction, nil)
	if err != nil {
		return RespLocationInfo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

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

// func (c *Client) Search() (Pokemon, error) {
// 	if c.Token == "" {
// 		return Pokemon{}, fmt.Errorf("log in to explore the world")
// 	}

// 	req, err := http.NewRequest("GET", c.baseURL+"/location/search", nil)
// 	if err != nil {
// 		return Pokemon{}, err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+c.Token)

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return Pokemon{}, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		err = bodyToError(resp.Body)
// 		return Pokemon{}, err
// 	}

// 	var response Pokemon
// 	err = json.NewDecoder(resp.Body).Decode(&response)
// 	if err != nil {
// 		return Pokemon{}, err
// 	}

// 	return response, nil
// }

func (c *Client) LocationSearch() (*websocket.Conn, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("log in to explore the world")
	}

	header := http.Header{}
	header.Add("Authorization", "Bearer "+c.Token)

	url := fmt.Sprintf("%s/location/search", c.baseURL)
	if strings.HasPrefix(c.baseURL, "http://") {
		url = strings.Replace(url, "http://", "ws://", 1)
	} else if strings.HasPrefix(c.baseURL, "https://") {
		url = strings.Replace(url, "https://", "wss://", 1)
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
