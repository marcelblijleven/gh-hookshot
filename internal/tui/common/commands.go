package common

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcelblijleven/gh-hookshot/internal/api"
)

func FetchWebhooksCmd(owner, repo string) tea.Cmd {
	return func() tea.Msg {
		var resp []api.Webhook

		if err := api.GetWebhooks(owner, repo, &resp); err != nil {
			return FetchWebhooksMsg{Err: err}
		}

		return FetchWebhooksMsg{Webhooks: resp}
	}
}

func FetchWebhookDeliveryDetailCmd(owner, repo string, hookID, deliveryID int) tea.Cmd {
	return func() tea.Msg {
		var resp api.HookDeliveryDetail

		if deliveryID == 0 {
			return nil
		}

		if err := api.GetWebhookDeliveryDetail(owner, repo, hookID, deliveryID, &resp); err != nil {
			return FetchDeliveryDetailMsg{Err: err}
		}

		return FetchDeliveryDetailMsg{DeliveryDetail: resp}
	}
}

func FetchWebhookDeliveriesCmd(owner, repo string, hookID int) tea.Cmd {
	return func() tea.Msg {
		var resp []api.HookDelivery

		if err := api.GetWebhookDeliveries(owner, repo, hookID, &resp); err != nil {
			return FetchDeliveriesMsg{Err: err}
		}

		return FetchDeliveriesMsg{Deliveries: resp}
	}
}

func FetchWebhookDeliveriesTickCmd(owner, repo string, hookID int) tea.Cmd {
	return tea.Tick(time.Second*5, func(t time.Time) tea.Msg {
		var resp []api.HookDelivery

		if err := api.GetWebhookDeliveries(owner, repo, hookID, &resp); err != nil {
			return FetchDeliveriesMsg{Err: err}
		}

		return FetchDeliveriesMsg{Deliveries: resp}
	})
}

func RedeliverWebhookDeliveryCmd(owner, repo string, hookID, deliveryID int) tea.Cmd {
	return func() tea.Msg {
		if deliveryID == 0 {
			return nil
		}
		if err := api.RedeliverWebhookDelivery(owner, repo, hookID, deliveryID); err != nil {
			return RedeliveryMsg{Err: err}
		}

		return RedeliveryMsg{}
	}
}

func SetNoDeliveriesCmd(hookID int) tea.Cmd {
	return func() tea.Msg {
		return NoDeliveriesMsg{HookID: hookID}
	}
}
