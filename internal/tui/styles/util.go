package styles

import "github.com/charmbracelet/lipgloss"

// GetTitleStyle returns the active or inactive title style based on
// the provided argument
func GetTitleStyle(active bool) lipgloss.Style {
	if active {
		return ActiveTitle
	}

	return InactiveTitle
}
