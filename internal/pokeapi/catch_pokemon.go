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
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
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
