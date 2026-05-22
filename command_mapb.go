package main

import (
	"fmt"

	"github.com/hardiing/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *Config) error {
	url := ""
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := pokeapi.GetLocationAreas(url, cfg.cache)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}
