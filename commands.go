package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Commands struct {
	name        string
	description string
	callback    func(*model, ...string) (tea.Cmd, error)
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
		"build": {
			name:        "build",
			description: "Build a <Building>",
			callback:    commandBuild,
		},
	}
	return variable
}
