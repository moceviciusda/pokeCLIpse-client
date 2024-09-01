package serverapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Register registers a new user with the given username and password
// and logs in the user.
func (c *Client) Register(username, password string) (RespLogin, error) {
	if c.Token != "" {
		return RespLogin{}, fmt.Errorf("already logged in")
	}

	type body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	payload, err := json.Marshal(body{Username: username, Password: password})
	if err != nil {
		return RespLogin{}, err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/users", bytes.NewReader(payload))
	if err != nil {
		return RespLogin{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLogin{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		err = bodyToError(resp.Body)
		return RespLogin{}, err
	}

	return c.Login(username, password)
}
