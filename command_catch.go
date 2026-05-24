package main

import (
	"fmt"
	"math/rand"

	"github.com/hardiing/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Missing parameters")
	}

	resp, err := pokeapi.GetPokemon(args[0], cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	exp := resp.BaseExperience
	chance := catchChance(exp)
	roll := rand.Intn(100)
	if roll < chance {
		fmt.Printf("%s was caught!\n", resp.Name)
		cfg.pokedex[resp.Name] = resp
	} else {
		fmt.Printf("%s escaped!\n", resp.Name)
	}

	return nil
}

func catchChance(exp int) int {
	if exp > 0 && exp <= 140 {
		return 90
	} else if exp > 140 && exp <= 280 {
		return 80
	} else if exp > 280 && exp <= 420 {
		return 70
	} else if exp > 420 && exp <= 560 {
		return 60
	} else if exp > 560 && exp <= 700 {
		return 40
	}
	return 0
}
