package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	leftBox := lipgloss.NewStyle().
		Width(30).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1).
		Render(fmt.Sprintf("Holz: %d\nStein: %d\nBevölkerung: %d", m.res.Wood, m.res.Stone, m.pop))

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
