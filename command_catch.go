package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	PokemonResp, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing ball at %s...\n", PokemonResp.Name)

	const threshold = 50
	catchLuck := rand.Intn(PokemonResp.BaseExperience)
	
	if threshold < catchLuck {
		return fmt.Errorf("%s escaped", PokemonResp.Name)
	}
	cfg.caughtPokemon[pokemonName] = PokemonResp

	fmt.Printf("%s was caught!\n", PokemonResp.Name)
	return nil
}