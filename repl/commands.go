package repl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type mapAPIResponse struct {
	Count int `json:"count"`
	config
	Results []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of locations",
			callback:    commandMap,
		},
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if c.Next != "" {
		url = c.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from the API: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", res.StatusCode)
	}

	var apiResponse mapAPIResponse
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return fmt.Errorf("failed to decode API response: %v", err)
	}

	for _, location := range apiResponse.Results {
		fmt.Println(location.Name)
	}

	*c = apiResponse.config

	return nil
}
