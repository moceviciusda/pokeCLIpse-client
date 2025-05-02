package serverapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
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

func (c *Client) SelectStarter() (*websocket.Conn, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("log in to explore the world")
	}

	header := http.Header{}
	header.Add("Authorization", "Bearer "+c.Token)

	url := fmt.Sprintf("%s/starter", c.baseURL)
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
