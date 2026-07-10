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
	config := config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		inputs := cleanInput(input)
		if len(inputs) == 0 {
			continue
		}

		commandName := inputs[0]
		if command, exists := getCommands()[commandName]; exists {
			err := command.callback(&config)
			if err != nil {
				fmt.Printf("Error occurred while executing command '%s': %v\n", command.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
