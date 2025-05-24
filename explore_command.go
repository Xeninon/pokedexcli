package main

import (
	"encoding/json"
	"fmt"

	"github.com/Xeninon/pokedexcli/internal/pokeapi"
	"github.com/Xeninon/pokedexcli/internal/pokecache"
)

type AreaPokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *Config, cache *pokecache.Cache, p map[string]Pokeinfo, area string) error {
	if area == "" {
		fmt.Println("area not specified")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + area

	body, err := pokeapi.PokeGet(cache, url)
	if err != nil {
		return err
	}

	areaPokemon := AreaPokemon{}
	err = json.Unmarshal(body, &areaPokemon)
	if err != nil {
		return err
	}

	for _, pokemon := range areaPokemon.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
