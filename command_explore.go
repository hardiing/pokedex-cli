package main

import (
	"fmt"

	"github.com/hardiing/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Missing parameters")
	}
	resp, err := pokeapi.GetArea(args[0], cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Name)

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	return err
}
