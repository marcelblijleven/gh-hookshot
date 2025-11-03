package deliverydetail

import "github.com/marcelblijleven/gh-hookshot/internal/api"

type HookDeliveryDetailItem struct {
	api.HookDeliveryDetail
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
