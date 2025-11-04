package webhooks

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcelblijleven/gh-hookshot/internal/api"
)

func FetchWebhooksCmd(owner, repo string) tea.Cmd {
	return func() tea.Msg {
		var resp []WebhookItem

		if err := api.GetWebhooks(owner, repo, &resp); err != nil {
			return webhooksFetchMsg{Err: err}
		}

		return webhooksFetchMsg{Webhooks: resp}
	}
}
