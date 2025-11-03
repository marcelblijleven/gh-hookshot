package deliveries

import (
	"fmt"

	"github.com/marcelblijleven/gh-hookshot/internal/api"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/styles"
)

type hookDeliveryItem struct {
	api.HookDelivery
}

// FilterValue satisfies list.Item
func (i hookDeliveryItem) FilterValue() string {
	return i.Event
}

// Title satisfies list.DetailItem
func (i hookDeliveryItem) Title() string {
	if i.Redelivery {
		return fmt.Sprintf("%s %s", i.DeliveredAt, styles.Repeat)
	}
	return i.DeliveredAt
}

// Description satisfies list.DetailItem
func (i hookDeliveryItem) Description() string {
	return fmt.Sprintf("Event: %s", i.Event)
}
