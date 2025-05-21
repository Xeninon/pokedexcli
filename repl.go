package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func pokerepl() {
	config := Config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		clean_text := cleanInput(text)
		command, ok := getCommands()[clean_text[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback(&config)
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
	callback    func(*Config) error
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
			description: "displays  previous 20 location areas",
			callback:    commandMapb,
		},
	}
}
