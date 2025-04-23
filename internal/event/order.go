package event

type OrderEvents struct {
	OrderCreated *OrderCreated
	OrderUpdated *OrderUpdated
	OrderListed  *OrderListed
}
