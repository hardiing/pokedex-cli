package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hardiing/pokedexcli/internal/pokecache"
)

type AreaDetails struct {
	Name              string       `json:"name"`
	PokemonEncounters []Encounters `json:"pokemon_encounters"`
}

type Encounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

func GetArea(areaInput string, cache *pokecache.Cache) (AreaDetails, error) {
	fullURL := "https://pokeapi.co/api/v2/location-area/" + areaInput
	checkCache, ok := cache.Get(fullURL)
	if ok {
		fmt.Println("cache used")
		jsonData := checkCache
		var area AreaDetails
		if err := json.Unmarshal(jsonData, &area); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}

		return area, nil
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
	var area AreaDetails
	if err := json.Unmarshal(jsonData, &area); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return area, err
}
