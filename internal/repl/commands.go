package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/erlanggajuni45/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, cache *pokecache.Cache, opts ...string) error
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

type exploreAPIResponse struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon locationArea `json:"pokemon"`
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
		"mapb": {
			name:        "mapb",
			description: "Displays a list of locations (backward)",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location and show list of Pokémon encounters",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a specific Pokémon",
			callback:    commandCatch,
		},
	}
}

func commandExit(c *config, cache *pokecache.Cache, opts ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, cache *pokecache.Cache, opts ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(c *config, cache *pokecache.Cache, opts ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if c.Next != "" {
		url = c.Next
	}

	var apiResponse mapAPIResponse

	// check if the response is already cached
	var data []byte
	if cachedData, exists := cache.Get(url); exists {
		data = cachedData
	} else {
		// fetch data from the API if not cached
		dat, err := fetchData(url, cache)
		if err != nil {
			return fmt.Errorf("failed to fetch data: %v", err)
		}
		data = dat
	}

	err := json.Unmarshal(data, &apiResponse)
	if err != nil {
		return fmt.Errorf("failed to decode API response: %v", err)
	}

	for _, location := range apiResponse.Results {
		fmt.Println(location.Name)
	}

	*c = apiResponse.config

	return nil
}

func commandMapb(c *config, cache *pokecache.Cache, opts ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if c.Previous != "" {
		url = c.Previous
	}

	var apiResponse mapAPIResponse

	// check if the response is already cached
	var data []byte
	if cachedData, exists := cache.Get(url); exists {
		data = cachedData
	} else {
		// fetch data from the API if not cached
		dat, err := fetchData(url, cache)
		if err != nil {
			return fmt.Errorf("failed to fetch data: %v", err)
		}
		data = dat
	}

	err := json.Unmarshal(data, &apiResponse)
	if err != nil {
		return fmt.Errorf("failed to decode API response: %v", err)
	}

	for _, location := range apiResponse.Results {
		fmt.Println(location.Name)
	}

	*c = apiResponse.config

	return nil
}

func commandExplore(c *config, cache *pokecache.Cache, opts ...string) error {
	location := opts[0]
	fmt.Println("Exploring location:", location)

	url := "https://pokeapi.co/api/v2/location-area/" + location

	var apiResponse exploreAPIResponse

	// check if the response is already cached
	var data []byte
	if cachedData, exists := cache.Get(url); exists {
		data = cachedData
	} else {
		// fetch data from the API if not cached
		dat, err := fetchData(url, cache)
		if err != nil {
			return fmt.Errorf("failed to fetch data: %v", err)
		}
		data = dat
	}

	err := json.Unmarshal(data, &apiResponse)
	if err != nil {
		return fmt.Errorf("failed to decode API response: %v", err)
	}

	if len(apiResponse.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon encounters found in this location.")
		return nil
	}

	fmt.Println("Found Pokémon:")

	for _, encounter := range apiResponse.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *config, cache *pokecache.Cache, opts ...string) error {
	return nil
}

func fetchData(url string, cache *pokecache.Cache) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from the API: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response body: %v", err)
	}

	cache.Add(url, data)

	return data, nil
}
