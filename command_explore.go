package main

import (
	"fmt"
	"errors"
)

func callbackExplore(cfg *config, args ...string) error{
	if len(args) < 1 {
		return errors.New("no pokeLocation provided")
	}
	pokeLocation := args[0]
	res, err := cfg.pokeapiClient.ListLocationData(pokeLocation)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", res.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}