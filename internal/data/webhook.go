package data

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cli/go-gh/v2/pkg/api"
)

type Webhook struct {
	Type          string       `json:"type"`
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Active        bool         `json:"active"`
	Events        []string     `json:"events"`
	Config        HookConfig   `json:"config"`
	UpdatedAt     string       `json:"updated_at"`
	CreatedAt     string       `json:"created_at"`
	URL           string       `json:"url"`
	TestURL       string       `json:"test_url"`
	PingURL       string       `json:"ping_url"`
	DeliveriesURL string       `json:"deliveries_url"`
	LastResponse  HookResponse `json:"last_response"`
}

type HookConfig struct {
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
	Secret      string `json:"secret"`
	InsecureSSL string `json:"insecure_ssl"`
}

type HookResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type HookDelivery struct {
	ID             int     `json:"id"`
	GUID           string  `json:"guid"`
	DeliveredAt    string  `json:"delivered_at"`
	Redelivery     bool    `json:"redelivery"`
	Duration       float64 `json:"duration"`
	StatusCode     int     `json:"status_code"`
	Event          string  `json:"event"`
	Action         string  `json:"action"`
	InstallationID int     `json:"installation_id"`
	RepositoryID   int     `json:"repository_id"`
	ThrottledAt    string  `json:"throttledAt"`
}

type HookDeliveryDetail struct {
	HookDelivery
	Request  HookDeliveryRequest  `json:"request"`
	Response HookDeliveryResponse `json:"response"`
}

type (
	HookDeliveryRequest struct {
		Payload json.RawMessage `json:"payload"`
	}
	HookDeliveryResponse struct {
		Payload string `json:"payload"`
	}
)

func GetWebhooks(owner, repo string, resp interface{}) error {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return err
	}

	return client.Get(fmt.Sprintf("repos/%s/%s/hooks", owner, repo), &resp)
}

func GetWebhookDeliveries(owner, repo string, hookID int, resp interface{}) error {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return err
	}

	idStr := strconv.Itoa(hookID)

	return client.Get(fmt.Sprintf("repos/%s/%s/hooks/%s/deliveries", owner, repo, idStr), &resp)
}

func GetWebhookDeliveryDetail(owner, repo string, hookID, deliveryID int, resp interface{}) error {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return err
	}

	idStr := strconv.Itoa(hookID)
	deliveryIDStr := strconv.Itoa(deliveryID)

	return client.Get(fmt.Sprintf("repos/%s/%s/hooks/%s/deliveries/%s", owner, repo, idStr, deliveryIDStr), &resp)
}
