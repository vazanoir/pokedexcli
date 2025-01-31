package repl

import (
	"fmt"
	"github.com/vazanoir/pokedexcli/pokeapi"
	"os"
)

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
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Show locations",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Go backwards in locations",
			callback:    mapbCommand,
		},
	}
}

func exitCommand(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(cfg *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func mapCommand(cfg *Config) error {
	current := cfg.Next
	page, err := pokeapi.GetLocations(current)
	if err != nil {
		return err
	}

	for _, loca := range page.Results {
		fmt.Println(loca.Name)
	}

	cfg.Prev = page.Previous
	cfg.Next = page.Next

	return nil
}

func mapbCommand(cfg *Config) error {
	current := cfg.Prev
	if current == "" {
		fmt.Println("You are on the first page")
		return nil
	}

	page, err := pokeapi.GetLocations(current)
	if err != nil {
		return err
	}

	for _, loca := range page.Results {
		fmt.Println(loca.Name)
	}

	cfg.Prev = page.Previous
	cfg.Next = page.Next

	return nil
}
