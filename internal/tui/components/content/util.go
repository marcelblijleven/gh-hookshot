package content

import "github.com/charmbracelet/lipgloss"

func getTitleStyle(column, current int) lipgloss.Style {
	if column == current {
		return activeTitleStyle
	}

	return inactiveTitleStyle
}
