package main

import (
	"time"

	"github.com/lazyturtlez/REPL-Pokedex/internal/pokeapi"
)

type config struct{
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
	caughtPokemon map[string]pokeapi.PokemonResp
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.PokemonResp),
		
	}
	startRepl(&cfg)
}
