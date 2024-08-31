package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/moceviciusda/pokeCLIpse-client/internal/serverapi"
)

type config struct {
	apiClient serverapi.Client
}

func startCLI(cfg *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("PokeCLIpse > ")
		scanner.Scan()

		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		inputWords := strings.Fields(input)

		command, ok := commands[inputWords[0]]
		if !ok {
			fmt.Println("invalid command:", input, "\ntype 'help' to see all available commands")
			continue
		}

		err := command.callback(cfg, inputWords[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
