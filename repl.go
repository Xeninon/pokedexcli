package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Xeninon/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func pokerepl() {
	config := Config{}
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		clean_text := cleanInput(text)
		if len(clean_text) == 0 {
			continue
		}
		param := ""
		if len(clean_text) > 1 {
			param = clean_text[1]
		}
		command, ok := getCommands()[clean_text[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback(&config, cache, param)
		}
	}
}

type Config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *pokecache.Cache, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "displays pokemon in a given location area (explore location_area)",
			callback:    commandExplore,
		},
	}
}
