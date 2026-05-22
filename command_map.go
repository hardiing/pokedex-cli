package main

import (
	"fmt"

	"github.com/hardiing/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	url := ""
	if cfg.Next != nil {
		url = *cfg.Next
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
