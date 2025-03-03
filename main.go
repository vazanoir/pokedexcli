package main

import (
	"time"

	"github.com/vazanoir/pokedexcli/internal/client"
	"github.com/vazanoir/pokedexcli/internal/repl"
)

func main() {
	c := client.NewClient(5*time.Second, 5*time.Minute)
	repl.StartRepl(&repl.Config{
		Next:   "https://pokeapi.co/api/v2/location-area",
		Client: c,
	})
}
