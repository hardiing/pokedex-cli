package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hardiing/pokedexcli/internal/pokeapi"
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

	commandHelp := func() error {
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
				callback:    commandExit(cfg),
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp(cfg),
			},
			"map": {
				name:        "map",
				description: "Display map areas",
				callback:    pokeapi.GetLocationAreas(cfg.Next),
			},
			/* "mapb": {
				name: "mapb",
				description: "Go back one page",
				callback: pokeapi.GetLocationAreas,
			}, */
		}
		c, ok := supportedCommands[r[0]]
		if ok {
			c.callback()
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
