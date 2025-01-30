package repl

import (
    "strings"
)

func cleanInput(text string) []string {
    lower_text := strings.ToLower(text)
    return strings.Fields(lower_text)
}
