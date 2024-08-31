package serverapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Login logs in the user with the given username and password
// and stores the session token in the client for future requests.
func (c *Client) Login(username, password string) (RespLogin, error) {
	type body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	payload, err := json.Marshal(body{Username: username, Password: password})
	if err != nil {
		return RespLogin{}, err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/login", bytes.NewReader(payload))
	if err != nil {
		return RespLogin{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLogin{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = bodyToError(resp.Body)
		return RespLogin{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLogin{}, err
	}

	response := RespLogin{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return RespLogin{}, err
	}

	c.token = response.Token
	return response, nil
}
