package main

import "fmt"

func commandHelp(cfg *config, params ...string) error {
	fmt.Println("Wecome to the PokeCLIpse!\nUsage:")
	fmt.Println()

	for key, command := range getCommands() {
		_, err := fmt.Printf("%v: %v\n", key, command.description)
		fmt.Println()
		if err != nil {
			return err
		}
	}

	return nil
}
