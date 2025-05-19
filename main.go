package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		clean_text := cleanInput(text)
		fmt.Printf("Your command was: %v\n", clean_text[0])
	}
}
