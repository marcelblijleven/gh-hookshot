package webhooks

import (
	"fmt"
	"strings"

	"github.com/marcelblijleven/gh-hookshot/internal/api"
)

type WebhookItem struct {
	api.Webhook
}

// FilterValue satisfies list.Item
func (i WebhookItem) FilterValue() string {
	return fmt.Sprintf("%s|%s", i.Config.URL, strings.Join(i.Events, "|"))
}

// Title satisfies list.DetailItem
func (i WebhookItem) Title() string {
	return i.Config.URL
}

// Description satisfies list.DetailItem
func (i WebhookItem) Description() string {
	var (
		events       string
		lastDelivery string
	)

	if len(i.Events) > 3 {
		events = fmt.Sprintf("Events: %d", len(i.Events))
	} else {
		events = fmt.Sprintf("Events: %s", strings.Join(i.Events, ", "))
	}

	lastDelivery = fmt.Sprintf("Last delivery: %s", i.LastResponse.Status)

	return fmt.Sprintf("%s\n%s", events, lastDelivery)
}
