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
}

func GetLocationAreas() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
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
	var areas []LocationName
	if err := json.Unmarshal(jsonData, &areas); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, area := range areas {
		fmt.Printf("%s", area.Name)
	}

	return nil
}
