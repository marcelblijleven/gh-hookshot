package common

import "github.com/marcelblijleven/gh-hookshot/internal/api"

type FetchWebhooksMsg struct {
	Webhooks []api.Webhook
	Err      error
}

type FetchDeliveriesMsg struct {
	Deliveries []api.HookDelivery
	Err        error
}

type FetchDeliveryDetailMsg struct {
	DeliveryDetail api.HookDeliveryDetail
	Err            error
}

type RedeliveryMsg struct {
	Err error
}

type NoDeliveriesMsg struct {
	HookID int
}
