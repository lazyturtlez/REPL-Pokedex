package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("unknown command")
			continue
		}
		
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 location areas in Pokemon world",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations in Pokemon world",
			callback: callbackMapB,
		},
		"explore": {
			name: "explore {location area}",
			description: "explores the provided poke location",
			callback:callbackExplore,
		},
		"catch": {
			name: "catch {pokemon name}",
			description: "attempts to catch a pokemon",
			callback:callbackCatch,
		},
		"inspect": {
			name: "inspect {pokemon name}",
			description: "inspects caught pokemon",
			callback:callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "inspects pokedex inventory",
			callback:callbackPokedex,
		},
	}
}