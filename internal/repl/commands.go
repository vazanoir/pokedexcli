package repl

import (
	"fmt"
	"os"
)

type cliCommand struct {
	Name     string
	Desc     string
	Callback func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name: "help",
			Desc: "Displays a help message",
			Callback: func(cfg *Config) error {
				fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

				for _, cmd := range GetCommands() {
					fmt.Printf("%s: %s\n", cmd.Name, cmd.Desc)
				}

				return nil
			},
		},
		"exit": {
			Name: "exit",
			Desc: "Exit the Pokedex",
			Callback: func(cfg *Config) error {
				fmt.Println("Closing the Pokedex... Goodbye!")
				os.Exit(0)
				return nil
			},
		},
		"map": {
			Name: "map",
			Desc: "Show locations",
			Callback: func(cfg *Config) error {
				current := cfg.Next
				page, err := cfg.Client.GetLocations(current)
				if err != nil {
					return err
				}

				for _, loca := range page.Results {
					fmt.Println(loca.Name)
				}

				cfg.Prev = page.Previous
				cfg.Next = page.Next

				return nil
			},
		},
		"mapb": {
			Name: "mapb",
			Desc: "Go backwards in locations",
			Callback: func(cfg *Config) error {
				current := cfg.Prev
				if current == "" {
					fmt.Println("You are on the first page")
					return nil
				}

				page, err := cfg.Client.GetLocations(current)
				if err != nil {
					return err
				}

				for _, loca := range page.Results {
					fmt.Println(loca.Name)
				}

				cfg.Prev = page.Previous
				cfg.Next = page.Next

				return nil
			},
		},
	}
}
