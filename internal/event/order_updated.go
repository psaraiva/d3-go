package event

import "time"

type OrderUpdated struct {
	Name    string
	Payload interface{}
}

func NewOrderUpdated() *OrderUpdated {
	return &OrderUpdated{
		Name: "OrderUpdated",
	}
}

func (e *OrderUpdated) GetName() string {
	return e.Name
}

func (e *OrderUpdated) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderUpdated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderUpdated) GetDateTime() time.Time {
	return time.Now()
}
