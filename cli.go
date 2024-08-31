package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
	"github.com/moceviciusda/pokeCLIpse-client/internal/serverapi"
)

type config struct {
	apiClient serverapi.Client
}

type completer struct {
	commands map[string]cliCommand
}

func (c *completer) Do(line []rune, pos int) (newLine [][]rune, length int) {
	words := strings.Fields(string(line))

	if len(words) <= 1 {
		// Autocomplete command names
		for cmd := range c.commands {
			if strings.HasPrefix(cmd, strings.ToLower(string(line))) {
				newLine = append(newLine, []rune(cmd[pos:]))
			}
		}
		return newLine, pos
	}

	// Autocomplete command parameters
	command, ok := c.commands[strings.ToLower(words[0])]
	if ok {
		for _, param := range command.params {
			lastWord := words[len(words)-1]
			if strings.HasPrefix(param, lastWord) {
				newLine = append(newLine, []rune(param[len(lastWord):]))
			}
		}
	}

	return newLine, pos
}

func startCLI(cfg *config) {
	commands := getCommands()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "PokeCLIpse > ",
		HistoryFile:     "/tmp/pokeclipse_history.txt",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    &completer{commands},
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		rl.SaveHistory(line)

		inputWords := strings.Fields(line)
		command, ok := commands[strings.ToLower(inputWords[0])]
		if !ok {
			fmt.Println("invalid command:", line, "\ntype 'help' to see all available commands")
			continue
		}

		err = command.callback(cfg, inputWords[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
