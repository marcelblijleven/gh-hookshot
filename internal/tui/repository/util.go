package repository

import (
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/styles"
)

func getRepoIcon(m Model) string {
	iconStyle := lipgloss.NewStyle()
	var icon string

	if m.err != nil || !m.repo.IsAdmin() {
		iconStyle = iconStyle.Foreground(styles.ColorWarning)
		icon = styles.CircleX
	} else {
		iconStyle = iconStyle.Foreground(styles.ColorGreen)
		icon = styles.CircleCheck
	}

	return iconStyle.Render(icon)
}
