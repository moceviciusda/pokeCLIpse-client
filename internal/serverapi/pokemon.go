package serverapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CreatePokemon(name string, level int, shiny bool) (Pokemon, error) {
	if c.token == "" {
		return Pokemon{}, fmt.Errorf("not logged in")
	}

	if level < 1 || level > 100 {
		return Pokemon{}, fmt.Errorf("invalid level. Must be between 1 and 100")
	}

	type parameters struct {
		Name  string `json:"name"`
		Level int    `json:"level"`
		Shiny bool   `json:"shiny"`
	}

	payload, err := json.Marshal(parameters{Name: name, Level: level, Shiny: shiny})
	if err != nil {
		return Pokemon{}, err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/pokemon", bytes.NewReader(payload))
	if err != nil {
		return Pokemon{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		err = bodyToError(resp.Body)
		return Pokemon{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	response := Pokemon{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return Pokemon{}, err
	}

	return response, nil
}
