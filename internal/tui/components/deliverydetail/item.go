package deliverydetail

import "github.com/marcelblijleven/gh-hookshot/internal/api"

type hookDeliveryDetailItem struct {
	api.HookDeliveryDetail
}

// FilterValue satisfies list.Item
func (i hookDeliveryDetailItem) FilterValue() string {
	return i.Event
}

// Title satisfies list.DetailItem
func (i hookDeliveryDetailItem) Title() string {
	return i.DeliveredAt
}

// Description satisfies list.DetailItem
func (i hookDeliveryDetailItem) Description() string {
	return i.GUID
}
