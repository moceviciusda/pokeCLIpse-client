package main

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, params ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Game",
			callback:    commandExit,
		},
	}
}
