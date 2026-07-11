package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/erlanggajuni45/pokedexcli/internal/pokecache"
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
	cache := pokecache.NewCache(10 * time.Second)
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
			err := command.callback(&config, &cache, inputs[1:]...)
			if err != nil {
				fmt.Printf("Error occurred while executing command '%s': %v\n", command.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
