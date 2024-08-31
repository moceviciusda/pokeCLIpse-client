package main

import "fmt"

func commandLogin(cfg *config, params ...string) error {
	if len(params) != 2 {
		return fmt.Errorf("login command requires 2 parameters: username and password\nUsage: login <username> <password>")
	}

	user, err := cfg.apiClient.Login(params[0], params[1])
	if err != nil {
		return err
	}

	fmt.Println("Welcome,", user.Username)

	return nil
}
