# Pokedex CLI

A simple command-line Pokédex application built with Go.  
This project interacts with the [PokéAPI](https://pokeapi.co/) to let users explore locations, catch Pokémon, and inspect captured Pokémon data from the terminal.

## Features

- Browse Pokémon locations
- Explore a location to see available Pokémon
- Catch Pokémon with a chance-based mechanic
- Inspect detailed information about caught Pokémon
- View your Pokédex collection
- Navigate forward and backward through location pages

## Requirements

- Go 1.20 or newer
- Internet connection to access PokéAPI

## Installation

Clone the repository:

```bash
git clone https://github.com/erlanggajuni45/pokedexcli.git
cd pokedexcli
```

Build the project:

```bash
go build
```

Run it:

```bash
./pokedexcli
```

Or run directly with:

```bash
go run .
```

## Usage

After starting the application, you can use commands such as:

- `help` — show available commands
- `map` — display the next page of locations
- `mapb` — display the previous page of locations
- `explore <location>` — list Pokémon found in a location
- `catch <pokemon>` — attempt to catch a Pokémon
- `inspect <pokemon>` — show details about a caught Pokémon
- `pokedex` — list all caught Pokémon
- `exit` — quit the application

## Example

```text
Pokedex > map
Pokedex > explore pastoria-city-area
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
```

## Project Structure

The project is organized as a small Go CLI application with logic for:

- command parsing
- API requests
- caching
- Pokémon data management
- terminal interaction

## Technologies Used

- Go
- PokéAPI
- Standard library packages

## Notes

- Catch success is randomized.
- Data is loaded from the public PokéAPI.
- Some responses may depend on network availability.

## License

This project is for learning purposes. Add a license if needed.