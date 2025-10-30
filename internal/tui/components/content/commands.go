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

func fetchWebhookDeliveriesCmd(owner, repo string, hookId int) tea.Cmd {
	return func() tea.Msg {
		var resp []HookDeliveryItem

		if err := data.GetWebhookDeliveries(owner, repo, hookId, &resp); err != nil {
			return deliveriesFetchMsg{Err: err}
		}

		return deliveriesFetchMsg{Deliveries: resp}
	}
}

func fetchWebhookDeliveryDetailCmd(owner, repo string, hookID, deliveryID int) tea.Cmd {
	return func() tea.Msg {
		var resp HookDeliveryDetailItem

		if err := data.GetWebhookDeliveryDetail(owner, repo, hookID, deliveryID, &resp); err != nil {
			return deliveryDetailFetchMsg{Err: err}
		}

		return deliveryDetailFetchMsg{DeliveryDetail: resp}
	}
}
