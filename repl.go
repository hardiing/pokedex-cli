package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"github.com/hardiing/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Previous *string
	Next     *string
}

func cleanInput(text string) []string {
	s := []string{}
	trimmedString := strings.Fields(text)
	for _, word := range trimmedString {
		lowerStr := strings.ToLower(word)
		s = append(s, lowerStr)
	}
	return s
}

func startRepl() {
	var supportedCommands map[string]cliCommand

	cfg := &Config{}

	commandHelp := func(cfg *Config) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("")
		for _, command := range supportedCommands {
			fmt.Printf("%s: %s\n", command.name, command.description)
		}
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		r := cleanInput((input))
		supportedCommands = map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Display the next 20 location areas",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Go back one page",
				callback:    commandMapb,
			},
		}
		c, ok := supportedCommands[r[0]]
		if ok {
			c.callback(cfg)
		} else {
			fmt.Println("Unknown command")
		}

	}
}
