package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const maxVisibleLogLines = 10

type model struct {
	input        string
	log          []string
	wood, stone  int
	pop          int
	scrollOffset int
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	m := model{
		wood:  10,
		stone: 5,
		pop:   3,
	}

	finalModel, err := tea.NewProgram(m).Run()
	if err != nil {
		fmt.Println("Fehler:", err)
		os.Exit(1)
	}

	if game, ok := finalModel.(model); ok {
		fmt.Println("Spiel beendet")
		fmt.Printf("Ressourcen: Holz: %d | Stein: %d | Bevölkerung: %d\n", game.wood, game.stone, game.pop)
	}
}
