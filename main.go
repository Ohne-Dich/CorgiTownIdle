package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	m := model{
		res: ResourceSet{
			Wood:  10,
			Stone: 5,
		},
		build: BuildingSet{
			Houses: 0,
		},
		pop:    3,
		popMax: 5,
		lang:   "ger",
	}

	finalModel, err := tea.NewProgram(m).Run()
	if err != nil {
		fmt.Println("Fehler:", err)
		os.Exit(1)
	}

	if game, ok := finalModel.(model); ok {
		fmt.Println("Spiel beendet")
		fmt.Printf("Ressourcen:\n")
		fmt.Printf("  Holz: %d\n", game.res.Wood)
		fmt.Printf("  Stein: %d\n", game.res.Stone)
		fmt.Printf("  Gold: %d\n", game.res.Gold)
		fmt.Printf("Gebäude:\n")
		fmt.Printf("  Häuser: %d\n", game.build.Houses)
		fmt.Printf("Bevölkerung: %d / %d\n", game.pop, game.popMax)
	}
}
