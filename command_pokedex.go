package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	if len(cfg.caughtPokemon) < 1 {
		return errors.New("pokedex is empty")
	}
	
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", pokemon)
	}
	return nil
}