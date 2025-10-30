package content

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcelblijleven/gh-hookshot/internal/data"
)

func fetchWebhooksCmd(owner, repo string) tea.Cmd {
	return func() tea.Msg {
		var resp []WebhookItem

		if err := data.GetWebhooks(owner, repo, &resp); err != nil {
			return webhooksFetchMsg{Err: err}
		}

		return webhooksFetchMsg{Webhooks: resp}
	}
}
