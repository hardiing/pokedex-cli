package main

import (
	"fmt"
)

func commandInspect(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Missing parameters")
	}

	exists, ok := cfg.pokedex[args[0]]
	if ok {
		fmt.Printf("Name: %s\n", exists.Name)
		fmt.Printf("Height: %d\n", exists.Height)
		fmt.Printf("Weight: %d\n", exists.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range exists.Stats {
			fmt.Printf("\t-%s: %d\n", stat.PokemonStat.StatName, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokemonType := range exists.Types {
			fmt.Printf("\t-%s\n", pokemonType.PokemonType.Name)
		}
	}
	return nil
}
