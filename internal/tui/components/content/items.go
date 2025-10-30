package content

import (
	"github.com/marcelblijleven/gh-hookshot/internal/data"
)

type WebhookItem struct {
	data.Webhook
}

type HookDeliveryItem struct {
	data.HookDelivery
}

type HookDeliveryDetailItem struct {
	data.HookDeliveryDetail
}

// FilterValue satisfies list.Item
func (i WebhookItem) FilterValue() string {
	return i.Config.URL
}

// Title satisfies list.DetailItem
func (i WebhookItem) Title() string {
	return i.Config.URL
}

// Description satisfies list.DetailItem
func (i WebhookItem) Description() string {
	return i.Config.URL
}

// FilterValue satisfies list.Item
func (i HookDeliveryItem) FilterValue() string {
	return i.Event
}

// Title satisfies list.DetailItem
func (i HookDeliveryItem) Title() string {
	return i.DeliveredAt
}

// Description satisfies list.DetailItem
func (i HookDeliveryItem) Description() string {
	return i.Event
}

// FilterValue satisfies list.Item
func (i HookDeliveryDetailItem) FilterValue() string {
	return i.Event
}

// Title satisfies list.DetailItem
func (i HookDeliveryDetailItem) Title() string {
	return i.DeliveredAt
}

// Description satisfies list.DetailItem
func (i HookDeliveryDetailItem) Description() string {
	return i.GUID
}
