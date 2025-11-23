package status

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	statusMsg     string
	statusTickMsg time.Time
)

func ShowStatus(msg string) tea.Cmd {
	return func() tea.Msg {
		return statusMsg(msg)
	}
}

func statusTick() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return statusTickMsg(t)
	})
}
