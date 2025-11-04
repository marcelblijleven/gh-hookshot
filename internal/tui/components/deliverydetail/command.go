package deliverydetail

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcelblijleven/gh-hookshot/internal/api"
)

func FetchWebhookDeliveryDetailCmd(owner, repo string, hookID, deliveryID int) tea.Cmd {
	return func() tea.Msg {
		var resp hookDeliveryDetailItem

		if err := api.GetWebhookDeliveryDetail(owner, repo, hookID, deliveryID, &resp); err != nil {
			return deliveryDetailFetchMsg{Err: err}
		}

		return deliveryDetailFetchMsg{DeliveryDetail: resp}
	}
}
