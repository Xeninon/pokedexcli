package main

import "fmt"

func commandMap(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != "" {
		url = c.Next
	}

	locationAreas := LocationAreas{}
	err := pokeGet(url, &locationAreas)
	if err != nil {
		return err
	}

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	print(c.Previous)
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	url := c.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas := LocationAreas{}
	err := pokeGet(url, &locationAreas)
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
