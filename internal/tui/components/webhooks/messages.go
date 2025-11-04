package webhooks

type webhooksFetchMsg struct {
	Webhooks []WebhookItem
	Err      error
}
