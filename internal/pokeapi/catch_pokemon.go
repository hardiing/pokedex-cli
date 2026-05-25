package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hardiing/pokedexcli/internal/pokecache"
)

type PokemonDetails struct {
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stats struct {
	BaseStat    int         `json:"base_stat"`
	PokemonStat PokemonStat `json:"stat"`
}

type PokemonStat struct {
	StatName string `json:"name"`
	Url      string `json:"url"`
}

type Types struct {
	Slot        int         `json:"slot"`
	PokemonType PokemonType `json:"type"`
}

type PokemonType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetPokemon(pokemonInput string, cache *pokecache.Cache) (PokemonDetails, error) {
	fullURL := "https://pokeapi.co/api/v2/pokemon/" + pokemonInput
	checkCache, ok := cache.Get(fullURL)
	if ok {
		fmt.Println("cache used")
		jsonData := checkCache
		var pokemon PokemonDetails
		if err := json.Unmarshal(jsonData, &pokemon); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}

		return pokemon, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	cache.Add(fullURL, body)
	fmt.Println("cache not used")
	jsonData := body
	var pokemon PokemonDetails
	if err := json.Unmarshal(jsonData, &pokemon); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return pokemon, err
}
