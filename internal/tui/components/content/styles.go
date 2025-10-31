package content

import "github.com/charmbracelet/lipgloss"

var containerStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingRight(1)

var activeTitleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230")).
	Padding(0, 1)

var inactiveTitleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("60")).
	Foreground(lipgloss.Color("230")).
	Padding(0, 1)
