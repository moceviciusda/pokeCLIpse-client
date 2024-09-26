package main

import (
	"fmt"

	"github.com/moceviciusda/pokeCLIpse-client/pkg/pokeutils"
)

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
	fmt.Println()

	err = selectStarterLoop(cfg)
	if err != nil {
		return err
	}

	return nil
}

func selectStarterLoop(cfg *config) error {
	isShiny := pokeutils.IsShiny()

	for {
		starter := selectOption(cfg.readline, "Select your starter Pokemon:", pokeutils.Starters, func(s []string) {
			for i, pokemon := range s {
				if i%3 == 0 {
					fmt.Println()
				}
				typeIcon := pokeutils.StarterTypeMap[pokemon]
				fmt.Printf("		%d. %s%s", i+1, pokemon, typeIcon)
			}
		})

		pokemon, err := cfg.apiClient.CreatePokemon(starter, 5, isShiny)
		if err != nil {
			fmt.Println("Error creating Pokemon:", err)
			continue
		}

		if pokemon.Shiny {
			pokemon.Name += "*"
		}

		cfg.apiClient.Party = append(cfg.apiClient.Party, pokemon)

		fmt.Printf("You received a %s<lvl %d>\n", pokemon.Name, pokemon.Level)
		fmt.Println()
		fmt.Println(pokemon.Stats)
		fmt.Println()

		return nil
	}
}
