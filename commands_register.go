package main

import "fmt"

func commandRegister(cfg *config, params ...string) error {
	if len(params) != 2 {
		return fmt.Errorf("register command requires 2 parameters: username and password\nUsage: register <username> <password>")

	}

	username := params[0]
	password := params[1]

	response, err := cfg.apiClient.Register(username, password)
	if err != nil {
		return err
	}

	fmt.Println("Successfully registered and logged in as", response.Username)
	return nil
}
