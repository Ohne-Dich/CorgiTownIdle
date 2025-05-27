package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func commandExit(m *model, args ...string) (tea.Cmd, error) {
	m.log = append(m.log, "Spiel wird beendet ...")
	return tea.Quit, nil
}

func commandHelp(m *model, args ...string) (tea.Cmd, error) {
	m.log = append(m.log, language[m.lang]["these are available"])
	for _, cmd := range getCommands(m) {
		m.log = append(m.log, "%s: %s\n", cmd.name, cmd.description)
	}
	return nil, nil
}
