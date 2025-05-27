package main

import (
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
)

// test
func populationIdle(m *model) (tea.Cmd, error) {
	for i := 0; i < m.pop; i++ {
		vari := rand.Int() % m.pop
		switch vari {
		case 0:
			m.pop++
		case 1:
			m.res.Stone++
		case 2:
			m.res.Wood++
		case 3:
			m.res.Gold++
		default:
			return nil, nil
		}
	}
	return nil, nil
}
