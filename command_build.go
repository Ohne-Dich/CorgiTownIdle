package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func commandBuild(m *model, args ...string) (tea.Cmd, error) {
	if len(args) < 1 {
		m.log = append(m.log, "Was möchtest du bauen?")
		return nil, nil
	}
	item := strings.ToLower(args[0])

	switch item {
	case "house":
		if m.wood < 5 || m.stone < 2 {
			m.log = append(m.log, "Nicht genug Ressourcen (5 Holz, 2 Stein benötigt)")
			return nil, nil
		}
		m.wood -= 5
		m.stone -= 2
		m.pop += 1
		m.log = append(m.log, "Haus gebaut! (+1 Bevölkerungs-Limit)")
	}

	return nil, nil
}
