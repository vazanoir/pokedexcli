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
        fmt.Print("Pokedex > ")
        scanner.Scan()
        input := scanner.Text()
        words := cleanInput(input)
        if len(words) < 1 {
            fmt.Println("Please input a command")
            continue
        }
        command := words[0]
        fmt.Printf("Your command was: %s\n", command)
    }
}

func cleanInput(text string) []string {
    lower_text := strings.ToLower(text)
    return strings.Fields(lower_text)
}
