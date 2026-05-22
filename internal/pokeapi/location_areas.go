package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hardiing/pokedexcli/internal/pokecache"
)

type LocationName struct {
	Name string
	URL  string
}

type LocationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationName `json:"results"`
}

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreasResponse, error) {
	fullURL := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if url != "" {
		fullURL = url
	}
	checkCache, ok := cache.Get(fullURL)
	if ok {
		fmt.Println("cache used")
		jsonData := checkCache
		var areas LocationAreasResponse
		if err := json.Unmarshal(jsonData, &areas); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}

		return areas, nil
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
	var areas LocationAreasResponse
	if err := json.Unmarshal(jsonData, &areas); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return areas, err
}
