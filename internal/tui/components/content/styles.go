package content

import "github.com/charmbracelet/lipgloss"

var containerStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingRight(1)

var titleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230")).
	Padding(0, 1)
