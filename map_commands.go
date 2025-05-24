package main

import (
	"encoding/json"
	"fmt"

	"github.com/Xeninon/pokedexcli/internal/pokeapi"
	"github.com/Xeninon/pokedexcli/internal/pokecache"
)

type LocationAreas struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *Config, cache *pokecache.Cache, s string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != "" {
		url = c.Next
	}

	body, err := pokeapi.PokeGet(cache, url)
	if err != nil {
		return err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapb(c *Config, cache *pokecache.Cache, s string) error {
	url := c.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	body, err := pokeapi.PokeGet(cache, url)
	if err != nil {
		return err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
