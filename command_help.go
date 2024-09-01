package main

import "fmt"

func commandHelp(cfg *config, params ...string) error {
	fmt.Println("Wecome to the PokeCLIpse!")
	fmt.Println("This is a command line Pokemon game where you can explore the world and catch Pokemon.")
	fmt.Println("Here are the available commands:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Println(command.name + ":")
		fmt.Println(command.description)
		fmt.Println()
	}

	return nil
}
