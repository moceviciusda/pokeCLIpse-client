package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/moceviciusda/pokeCLIpse-client/internal/serverapi"
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

	for {
		err = selectStarter(cfg)
		if err != nil {
			fmt.Println("Error selecting starter:", err)
			continue
		}
		return nil
	}
}

func selectStarter(cfg *config) error {
	conn, err := cfg.apiClient.SelectStarter()
	if err != nil {
		return err
	}

	defer conn.Close()

	var message struct {
		Error   string   `json:"error"`
		Message string   `json:"message"`
		Options []string `json:"options"`
	}
	err = conn.ReadJSON(&message)
	if err != nil {
		return err
	}
	if message.Error != "" || len(message.Options) <= 0 {
		return fmt.Errorf(message.Error)
	}

	var selected string
	for {
		selected = selectOption(cfg.readline, "Select your starter Pokemon:", message.Options, func(s []string) {
			for i, pokemon := range s {
				if i%3 == 0 {
					fmt.Println()
				}
				typeIcon := pokeutils.StarterTypeMap[pokemon]
				fmt.Printf("\t\t%d. %s%s", i+1, pokemon, typeIcon)
			}
			fmt.Println()
		})
		if selected != "Interrupt" {
			break
		}
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte(selected))
	if err != nil {
		return err
	}

	pokemon := struct {
		serverapi.Pokemon
		Error string `json:"error"`
	}{}
	err = conn.ReadJSON(&pokemon)
	if err != nil {
		return err
	}
	if pokemon.Error != "" {
		return fmt.Errorf(pokemon.Error)
	}

	if pokemon.Shiny {
		pokemon.Name += "*"
	}

	fmt.Printf("You received a %s<lvl %d>\n", pokemon.Name, pokemon.Level)
	fmt.Println()
	fmt.Println(pokemon.Stats)
	fmt.Println()
	cfg.apiClient.Party = append(cfg.apiClient.Party, pokemon.Pokemon)
	return nil
}
