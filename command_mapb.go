package main

import (
	"fmt"
)

func callbackMapB(cfg *config, words ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return fmt.Errorf("currently on the first page")
	}

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous
	return nil
}