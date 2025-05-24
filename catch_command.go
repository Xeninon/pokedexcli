package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/Xeninon/pokedexcli/internal/pokeapi"
	"github.com/Xeninon/pokedexcli/internal/pokecache"
)

type Pokeinfo struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Name           string `json:"name"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func commandCatch(c *Config, cache *pokecache.Cache, pokedex map[string]Pokeinfo, pokemon string) error {
	if pokemon == "" {
		fmt.Println("pokemon not specified")
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	body, err := pokeapi.PokeGet(cache, url)
	if err != nil {
		return err
	}

	pokeinfo := Pokeinfo{}
	err = json.Unmarshal(body, &pokeinfo)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	if float32(rand.Intn(600))*rand.Float32() > float32(pokeinfo.BaseExperience) {
		fmt.Printf("%v was caught!\n", pokemon)
		pokedex[pokemon] = pokeinfo
		return nil
	}

	fmt.Printf("%v escaped!\n", pokemon)
	return nil
}
