package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[31mPokedex > \033[0m")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) < 1 {
			fmt.Println("Please input a command")
			continue
		}

		cmd, found := getCommands()[words[0]]
		if !found {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Printf("error using %s's callback: %v", cmd.name, err)
		}
	}
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	return strings.Fields(lower_text)
}
