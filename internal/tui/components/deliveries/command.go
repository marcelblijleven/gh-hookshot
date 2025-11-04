package deliveries

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcelblijleven/gh-hookshot/internal/api"
)

func FetchWebhookDeliveriesCmd(owner, repo string, hookId int) tea.Cmd {
	return func() tea.Msg {
		var resp []hookDeliveryItem

		if err := api.GetWebhookDeliveries(owner, repo, hookId, &resp); err != nil {
			return deliveriesFetchMsg{Err: err}
		}

		return deliveriesFetchMsg{Deliveries: resp}
	}
}
