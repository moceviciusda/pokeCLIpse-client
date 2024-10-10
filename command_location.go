package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/moceviciusda/pokeCLIpse-client/internal/serverapi"
)

func commandLocation(cfg *config, params ...string) error {
	if len(params) == 0 {
		location, err := cfg.apiClient.LocationInfo()
		if err != nil {
			return err
		}

		fmt.Println("You are at", location.Name)
		fmt.Println()
		// TODO: add some more information about the location

		fmt.Println("You can move using 'location next' or 'location previous' commands")
		fmt.Println("Previous:", location.Previous, "| Next:", location.Next)
		fmt.Println()

		return nil
	}

	switch params[0] {
	case "next", "previous":
		location, err := cfg.apiClient.Move(params[0])
		if err != nil {
			return err
		}

		fmt.Println("You are now at", location.Name)
		fmt.Println()
		fmt.Println("Previous:", location.Previous, "| Next:", location.Next)
		fmt.Println()
		return nil

	case "search":

		conn, err := cfg.apiClient.LocationSearch()
		if err != nil {
			return err
		}

		defer conn.Close()

		var pokemon struct {
			Error string `json:"error"`
			serverapi.Pokemon
		}
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

		fmt.Printf("You encountered a wild %s<lvl %d>\n", pokemon.Name, pokemon.Level)
		fmt.Println()
		fmt.Println(pokemon.Stats)
		fmt.Println()

		action := selectOption(cfg.readline, "What do you want to do?", []string{"battle", "run"}, nil)

		switch action {
		case "battle":
			cfg.handleWildBattle(conn)

		case "run":
			conn.WriteMessage(websocket.TextMessage, []byte("run"))
			fmt.Println("You ran away!")
			fmt.Println()
		default:
			return fmt.Errorf("invalid action. Use 'battle' or 'run'")
		}

		return nil

	default:
		return fmt.Errorf("invalid parameter. Use 'next' or 'previous' to move, or 'search' to find wild pokemon")
	}
}

func (cfg *config) handleWildBattle(conn *websocket.Conn) error {
	conn.WriteMessage(websocket.TextMessage, []byte("battle"))

	for {
		var message struct {
			Error   string   `json:"error"`
			Message string   `json:"message"`
			Options []string `json:"options"`
		}
		err := conn.ReadJSON(&message)
		if err != nil {
			return err
		}

		if message.Error != "" {
			return fmt.Errorf(message.Error)
		}

		if len(message.Options) > 0 {
			action := selectOption(cfg.readline, message.Message, message.Options, nil)

			err := conn.WriteMessage(websocket.TextMessage, []byte(action))
			if err != nil {
				return err
			}
			continue
		}

		fmt.Println(message.Message)
	}
}
