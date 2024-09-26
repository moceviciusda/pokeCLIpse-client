package main

import "github.com/moceviciusda/pokeCLIpse-client/pkg/ansiiutils"

type cliCommand struct {
	name        string
	description string
	params      []string
	callback    func(config *config, params ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        ansiiutils.StyleBold + "help" + ansiiutils.Reset,
			description: "	Displays a help message",
			callback:    commandHelp,
		},
		"register": {
			name:        ansiiutils.StyleBold + "register" + ansiiutils.Reset + ansiiutils.StyleItalic + " <username> <password>" + ansiiutils.Reset,
			description: "	Register a new account",
			callback:    commandRegister,
		},
		"login": {
			name:        ansiiutils.StyleBold + "login" + ansiiutils.Reset + ansiiutils.StyleItalic + " <username> <password>" + ansiiutils.Reset,
			description: "	Login to the game",
			callback:    commandLogin,
		},

		"location": {
			name: ansiiutils.StyleBold + "location" + ansiiutils.Reset + ansiiutils.StyleItalic + " [next|previous|search]" + ansiiutils.Reset,
			description: "	Get information about your current location\n" +
				"	Use 'next' or 'previous' to move to a different location\n" +
				"	Use 'search' to find wild Pokemon in the current location",
			params:   []string{"next", "previous", "search"},
			callback: commandLocation,
		},
		"exit": {
			name:        ansiiutils.StyleBold + "exit" + ansiiutils.Reset,
			description: "	Exit the Game",
			callback:    commandExit,
		},
	}
}
