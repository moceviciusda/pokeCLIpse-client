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

	party, err := cfg.apiClient.GetPokemonParty()
	if err != nil {
		return err
	}

	cfg.apiClient.Party = party

	fmt.Println("Welcome back,", user.Username)
	fmt.Println("You have the following Pokemon in your party:")
	fmt.Println()
	for i, pokemon := range party {
		fmt.Printf("	%d. %s<lvl %d>\n", i+1, pokemon.Name, pokemon.Level)
	}
	fmt.Println()

	return nil
}
