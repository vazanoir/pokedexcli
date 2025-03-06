package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vazanoir/pokedexcli/internal/client"
)

type Config struct {
	Client   client.Client
	Commands map[string]cliCommand
	Pokedex  map[string]client.Pokemon
	Prev     string
	Next     string
}

func StartRepl(cfg *Config) {
	cfg.Commands = InitCommands()
	cfg.Pokedex = make(map[string]client.Pokemon)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[31mPokedex > \033[0m")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) < 1 {
			fmt.Println("Please input a command")
			continue
		}

		cmd, found := cfg.Commands[words[0]]
		if !found {
			fmt.Println("Unknown command")
			continue
		}

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		err := cmd.Callback(cfg, args...)
		if err != nil {
			fmt.Printf("error using %s's callback: %v", cmd.Name, err)
		}
	}
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	return strings.Fields(lower_text)
}
