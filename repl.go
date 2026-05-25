package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hardiing/pokedexcli/internal/pokeapi"
	"github.com/hardiing/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args []string) error
}

type Config struct {
	Previous *string
	Next     *string
	cache    *pokecache.Cache
	pokedex  map[string]pokeapi.PokemonDetails
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
	cfg.cache = pokecache.NewCache(5 * time.Second)
	cfg.pokedex = make(map[string]pokeapi.PokemonDetails)

	commandHelp := func(cfg *Config, args []string) error {
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
		command := r[0]
		args := r[1:]
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
			"explore": {
				name:        "explore",
				description: "Explore an area by name",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "attempt to catch a pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "inspect a pokemon in your pokedex",
				callback:    commandInspect,
			},
		}
		c, ok := supportedCommands[command]
		if ok {
			c.callback(cfg, args)
		} else {
			fmt.Println("Unknown command")
		}

	}
}
