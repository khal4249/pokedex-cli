package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanLine := cleanInput(line)
		//fmt.Printf("Your command was: %s\n", cleanLine[0])
		command := cleanLine[0]
		commands := returnCommands()
		values, ok := commands[command]
		if ok {
			values.callback()
			//err := values.callback
			//fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
	/*if err := scanner.Err(); err != nil {
	    fmt.Fprintln(os.Stderr, "error:", err)
	} */
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	cleanedSlice := strings.Fields(cleaned)
	return cleanedSlice
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	commands := returnCommands()
	for command, desc := range commands {
		fmt.Printf("%s: %s\n", command, desc.description)
	}
	return nil
}

func returnCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	return commands
}
