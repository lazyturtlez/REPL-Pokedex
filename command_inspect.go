package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon provided")
	}
	pokemon := args[0]

	pokemonInfo, exists := cfg.caughtPokemon[pokemon]
	if !exists {
		return fmt.Errorf("you hva not caught %s", pokemon)
	}

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %v\n", pokemonInfo.Height)
	fmt.Printf("Weight: %v\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemonInfo.Types {
		fmt.Printf(" - %v\n", types.Type.Name)
	}
	return nil
}