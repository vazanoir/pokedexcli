package repl

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "",
            expected: []string{},
        },
        {
            input: "hello",
            expected: []string{"hello"},
        },
        {
            input: "  hello world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "  hello WORLD  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "  hello-WORLD  ",
            expected: []string{"hello-world"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)

        if len(actual) != len(c.expected) {
            t.Errorf("different output length: %v and %v", len(actual), len(c.expected))
            continue
        }

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]

            if word != expectedWord {
                t.Errorf("different words: '%s' and '%s'", word, expectedWord)
                continue
            }
        }
    }
}
