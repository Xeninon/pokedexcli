package main

import (
	"fmt"
	"os"

	"github.com/Xeninon/pokedexcli/internal/pokecache"
)

func commandHelp(c *Config, ca *pokecache.Cache, p map[string]Pokeinfo, s string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandExit(c *Config, ca *pokecache.Cache, p map[string]Pokeinfo, s string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandPokedex(c *Config, ca *pokecache.Cache, pokedex map[string]Pokeinfo, s string) error {
	fmt.Println("Your Pokedex:")
	for pokemon, _ := range pokedex {
		fmt.Println(" - ", pokemon)
	}

	return nil
}
