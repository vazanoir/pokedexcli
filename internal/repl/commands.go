package repl

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type cliCommand struct {
	Name     string
	Desc     string
	Callback func(*Config, ...string) error
}

func InitCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name: "help",
			Desc: "Displays a help message",
			Callback: func(cfg *Config, args ...string) error {
				fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

				commandsOrder := []string{
					"help",
					"exit",
					"map",
					"mapb",
					"explore",
					"catch",
				}

				for _, cmdName := range commandsOrder {
					cmd := cfg.Commands[cmdName]
					fmt.Printf("%s: %s\n", cmd.Name, cmd.Desc)
				}

				return nil
			},
		},
		"exit": {
			Name: "exit",
			Desc: "Exit the Pokedex",
			Callback: func(cfg *Config, args ...string) error {
				fmt.Println("Closing the Pokedex... Goodbye!")
				os.Exit(0)
				return nil
			},
		},
		"map": {
			Name: "map",
			Desc: "Show locations",
			Callback: func(cfg *Config, args ...string) error {
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
			Callback: func(cfg *Config, args ...string) error {
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
		"explore": {
			Name: "explore",
			Desc: "Give this command a location name to explore it!",
			Callback: func(cfg *Config, args ...string) error {
				if len(args) < 1 {
					fmt.Printf("Missing a location name, use map or mapb to find some.\n")
					return nil
				}
				arg := args[0]

				loc, err := cfg.Client.ExploreLocations("https://pokeapi.co/api/v2/location-area/" + arg)
				if err != nil {
					errParts := strings.Split(err.Error(), " ")
					errCode, err := strconv.Atoi(errParts[len(errParts)-1])
					if err != nil {
						return err
					}

					switch errCode {
					case 404:
						fmt.Printf("Location %v not found, are you lost? Use map or mapb.\n", arg)
					default:
						return err
					}

					return nil
				}

				fmt.Printf("Exploring %v...\n", loc.Name)
				fmt.Printf("Found Pokemon:\n")
				for _, poke := range loc.PokemonEncounters {
					fmt.Printf(" - %v\n", poke.Pokemon.Name)
				}

				return nil
			},
		},
		"catch": {
			Name: "catch",
			Desc: "Try to catch a pokemon!",
			Callback: func(cfg *Config, args ...string) error {
				if len(args) != 1 {
					fmt.Printf("Input one pokemon name, use explore to find some.\n")
					return nil
				}
				arg := args[0]

				pokemon, err := cfg.Client.CatchPokemon("https://pokeapi.co/api/v2/pokemon/" + arg)
				if err != nil {
					errParts := strings.Split(err.Error(), " ")
					errCode, err := strconv.Atoi(errParts[len(errParts)-1])
					if err != nil {
						return err
					}

					switch errCode {
					case 404:
						fmt.Printf("Pokemon %v not found. Use explore to find valid ones.\n", arg)
					default:
						return err
					}

					return nil
				}

				fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
				// max BaseExperience can be 635
				maxThrow := 650
				throw := rand.IntN(maxThrow)
				chance := (1 - float64(pokemon.BaseExperience) / float64(maxThrow)) * 100
				fmt.Printf("%.0f%% chance of capturing %v!\n", chance, pokemon.Name)
				if throw > pokemon.BaseExperience {
					cfg.Pokedex[pokemon.Name] = pokemon
					fmt.Printf("%v catched!\n", cases.Title(language.English, cases.NoLower).String(pokemon.Name))
				} else {
					fmt.Printf("%v escaped!\n", cases.Title(language.English, cases.NoLower).String(pokemon.Name))
				}

				return nil
			},
		},
	}
}
