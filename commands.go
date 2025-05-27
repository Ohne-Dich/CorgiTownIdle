package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Commands struct {
	name        string
	description string
	callback    func(*model) (tea.Cmd, error)
}

func getCommands(m *model) map[string]Commands {
	variable := map[string]Commands{
		"exit": {
			name:        "exit",
			description: "Exit the game",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Shows current commands",
			callback:    commandHelp,
		},
	}
	return variable
}
