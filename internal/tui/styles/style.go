package styles

import "github.com/charmbracelet/lipgloss"

const DeliveriesWidth = 30

var (
	ActiveTitle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1).
			MarginLeft(2)
	InactiveTitle = lipgloss.NewStyle().
			Background(lipgloss.Color("60")).
			Foreground(lipgloss.Color("250")).
			Padding(0, 1).
			MarginLeft(2)
	Container = lipgloss.NewStyle().Padding(0, 1, 0, 1)
)
