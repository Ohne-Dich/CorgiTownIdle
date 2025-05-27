package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.input != "" {
				m.log = append(m.log, m.input)
				m.input = ""
				m.scrollOffset = 0 // zurück zum neuesten Eintrag
			}
		case "up":
			if m.scrollOffset < len(m.log)-maxVisibleLogLines {
				m.scrollOffset++
			}
		case "down":
			if m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "ctrl+c":
			return m, tea.Quit
		default:
			m.input += msg.String()
		}
	}

	return m, nil
}

func (m model) View() string {
	leftBox := lipgloss.NewStyle().
		Width(30).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1).
		Render(fmt.Sprintf("Holz: %d\nStein: %d\nBevölkerung: %d", m.wood, m.stone, m.pop))

	rightBox := lipgloss.NewStyle().
		Width(30).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1).
		Render("Einheiten:\n")

	topRow := lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox)

	logStyle := lipgloss.NewStyle().
		Width(40).
		Height(maxVisibleLogLines+2). // +2 wegen Titel und Padding
		Padding(0, 1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("241"))

	logs := formatLog(m.log, m.scrollOffset)
	logView := logStyle.Render("Logs:\n" + logs)

	input := lipgloss.NewStyle().
		BorderTop(true).
		Padding(0, 1).
		Render("> " + m.input)

	return lipgloss.JoinVertical(lipgloss.Left, topRow, logView, input)
}

func formatLog(log []string, offset int) string {
	total := len(log)
	if total == 0 {
		return "(keine Aktionen)"
	}

	maxLines := maxVisibleLogLines
	if total < maxLines {
		maxLines = total
	}

	start := total - maxLines - offset
	if start < 0 {
		start = 0
	}
	end := start + maxLines
	if end > total {
		end = total
	}

	out := ""
	for _, l := range log[start:end] {
		out += "- " + l + "\n"
	}
	return out
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
