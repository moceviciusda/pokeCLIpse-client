package serverapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CreatePokemon(name string, level int, shiny bool) (Pokemon, error) {
	if c.Token == "" {
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

	req.Header.Set("Authorization", "Bearer "+c.Token)

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

func (c *Client) GetPokemonParty() ([]Pokemon, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("not logged in")
	}

	req, err := http.NewRequest("GET", c.baseURL+"/pokemon/party", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = bodyToError(resp.Body)
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := []Pokemon{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
