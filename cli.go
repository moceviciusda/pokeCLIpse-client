package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	"github.com/moceviciusda/pokeCLIpse-client/internal/serverapi"
)

type config struct {
	apiClient serverapi.Client
	readline  *readline.Instance
}

func startCLI(cfg *config) {
	commands := getCommands()
	completer := completer{commands: commands, selectOptions: []string{}}

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "PokeCLIpse > ",
		HistoryFile:     "/tmp/pokeclipse_history.txt",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    &completer,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	cfg.readline = rl

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

type completer struct {
	commands      map[string]cliCommand
	selectOptions []string
}

func (c *completer) Do(line []rune, pos int) (newLine [][]rune, length int) {
	if len(c.selectOptions) > 0 {
		for _, opt := range c.selectOptions {
			if strings.HasPrefix(opt, strings.ToLower(string(line))) {
				newLine = append(newLine, []rune(opt[pos:]))
			}
		}
		return newLine, pos
	}

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

func selectOption(rl *readline.Instance, prompt string, options []string) (string, error) {
	fmt.Println(prompt)
	for i, opt := range options {
		fmt.Printf("	%d. %s\n", i+1, opt)
	}

	rl.Config.AutoComplete.(*completer).selectOptions = options
	defer func() {
		rl.Config.AutoComplete.(*completer).selectOptions = []string{}
	}()

	for {
		line, err := rl.Readline()
		if err != nil {
			return "", err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if choice, err := strconv.Atoi(line); err == nil && choice > 0 && choice <= len(options) {
			return options[choice-1], nil
		}

		for i, opt := range options {
			if strings.EqualFold(line, opt) {
				return options[i], nil
			}
		}

		fmt.Println("Invalid option. Please try again.")
	}
}
