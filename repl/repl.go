package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")
}

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		inputs := cleanInput(input)
		if len(inputs) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", inputs[0])
	}
}
