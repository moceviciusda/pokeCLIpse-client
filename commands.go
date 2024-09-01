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
			description: "	Displays a help message",
			callback:    commandHelp,
		},
		"register": {
			name:        "register <username> <password>",
			description: "	Register a new account",
			callback:    commandRegister,
		},
		"login": {
			name:        "login <username> <password>",
			description: "	Login to the game",
			callback:    commandLogin,
		},

		"location": {
			name: "location [next|previous|search]",
			description: "	Get information about your current location\n" +
				"	Use 'next' or 'previous' to move to a different location\n" +
				"	Use 'search' to find wild Pokemon in the current location",
			params:   []string{"next", "previous", "search"},
			callback: commandLocation,
		},
		"exit": {
			name:        "exit",
			description: "	Exit the Game",
			callback:    commandExit,
		},
	}
}
