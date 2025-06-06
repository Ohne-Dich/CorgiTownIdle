package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			cmdText := strings.TrimSpace(m.input)
			m.input = ""
			populationIdle(&m)

			if cmdText != "" {
				m.log = append(m.log, "> "+cmdText)
				return handleCommand(m, cmdText)
			}
			return m, nil

		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			return m, nil

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		default:
			// Nur sichtbare Zeichen hinzufügen
			if msg.String() != "" && len(msg.String()) == 1 {
				m.input += msg.String()
			}
			return m, nil
		}
	}

	return m, nil
}

func handleCommand(m model, input string) (tea.Model, tea.Cmd) {
	words := strings.Fields(input)
	command := words[0]
	args := words[1:]
	commands := getCommands(&m)
	if cmd, ok := commands[strings.ToLower(command)]; ok {
		teaCmd, err := cmd.callback(&m, args...)
		if err != nil {
			m.log = append(m.log, "Fehler: "+err.Error())
		}
		return m, teaCmd
	}

	m.log = append(m.log, "Unbekannter Befehl: "+input)
	return m, nil
}
