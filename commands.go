package main

type cliCommand struct {
	name        string
	description string
	params      []string
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
		"login": {
			name:        "login <username> <password>",
			description: "Login to the game",
			callback:    commandLogin,
		},
		"register": {
			name:        "register <username> <password>",
			description: "Register a new account",
			callback:    commandRegister,
		},
		"location": {
			name:        "location [next|previous|search]",
			description: "Get information about the current location or use 'next' or 'previous' to move;\n use 'search' to find wild pokemon in the area",
			params:      []string{"next", "previous", "search"},
			callback:    commandLocation,
		},
	}
}
