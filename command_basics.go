package main

import tea "github.com/charmbracelet/bubbletea"

func commandExit(m *model) (tea.Cmd, error) {
	m.log = append(m.log, "Spiel wird beendet ...")
	return tea.Quit, nil
}

func commandHelp(m *model) (tea.Cmd, error) {
	m.log = append(m.log, "Verfügbare Befehle: exit, help …")
	return nil, nil
}
