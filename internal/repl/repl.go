package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vazanoir/pokedexcli/internal/client"
)

type Config struct {
	Client client.Client
	Prev   string
	Next   string
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[31mPokedex > \033[0m")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) < 1 {
			fmt.Println("Please input a command")
			continue
		}

		cmd, found := GetCommands()[words[0]]
		if !found {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.Callback(cfg)
		if err != nil {
			fmt.Printf("error using %s's callback: %v", cmd.Name, err)
		}
	}
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	return strings.Fields(lower_text)
}
