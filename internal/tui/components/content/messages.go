package content

type webhooksFetchMsg struct {
	Webhooks []WebhookItem
	Err      error
}

type deliveriesFetchMsg struct {
	Deliveries []HookDeliveryItem
	Err        error
}
