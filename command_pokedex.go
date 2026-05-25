package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, args []string) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.pokedex {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}
