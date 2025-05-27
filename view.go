package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	leftBox := lipgloss.NewStyle().
		Width(30).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1).
		Render(renderStructFields(m.res, m.lang))

	rightBox := lipgloss.NewStyle().
		Width(30).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1).
		Render(fmt.Sprintf("Einheiten:\nBewohner: %v\nMax: %v\n", m.pop, m.popMax))

	topRow := lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox)

	logStyle := lipgloss.NewStyle().
		Width(62).
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

func renderStructFields(data any, lang string) string {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	var sb strings.Builder
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		localized := language[lang][fieldName]
		if localized == "" {
			localized = fieldName // fallback
		}
		value := v.Field(i).Interface()
		sb.WriteString(fmt.Sprintf("%s: %v\n", localized, value))
	}
	return sb.String()
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
