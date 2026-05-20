package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func GetLocationAreas(url string) (LocationAreasResponse, error) {
	fullURL := "https://pokeapi.co/api/v2/location-area/"
	if url != "" {
		fullURL = url
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

	jsonData := body
	var areas LocationAreasResponse
	if err := json.Unmarshal(jsonData, &areas); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	//config.Next = res.Next

	return areas, err
}
