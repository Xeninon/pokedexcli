package main

import (
	"fmt"

	"github.com/Xeninon/pokedexcli/internal/pokecache"
)

func commandInspect(c *Config, ca *pokecache.Cache, pokedex map[string]Pokeinfo, pokemon string) error {
	if pokemon == "" {
		fmt.Println("pokemon not specified")
		return nil
	}

	pokeinfo, ok := pokedex[pokemon]
	if !ok {
		fmt.Println("pokemon not caught yet")
		return nil
	}

	fmt.Printf("Name: %v\n", pokeinfo.Name)
	fmt.Printf("Height: %v\n", pokeinfo.Height)
	fmt.Printf("Weight: %v\n", pokeinfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokeinfo.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, ptype := range pokeinfo.Types {
		fmt.Printf("  -%v\n", ptype.Type.Name)
	}

	return nil
}
