package deliveries

type deliveriesFetchMsg struct {
	Deliveries []hookDeliveryItem
	Err        error
}
