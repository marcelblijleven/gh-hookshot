package webhooks

import tea "github.com/charmbracelet/bubbletea"

func webhookSelectedCmd(hookID int) tea.Cmd {
	return func() tea.Msg {
		return WebhookSelectedMsg{HookID: hookID}
	}
}
