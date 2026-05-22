package main

import (
	"fmt"

	"github.com/hardiing/pokedexcli/internal/pokeapi"
)

func commandExplore(name string, cfg *Config) error {
	url := ""
	resp, err := pokeapi.GetLocationAreas(url, cfg.cache)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		if name == area.Name {
			fmt.Printf("Exploring %s...", area.Name)
			// explore logic
			// print pokemon
		}
	}

	return err
}
