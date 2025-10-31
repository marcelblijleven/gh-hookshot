package content

type webhooksFetchMsg struct {
	Webhooks []WebhookItem
	Err      error
}

type deliveriesFetchMsg struct {
	Deliveries []HookDeliveryItem
	Err        error
}

type deliveryDetailFetchMsg struct {
	DeliveryDetail HookDeliveryDetailItem
	Err            error
}
