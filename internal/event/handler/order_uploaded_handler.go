package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"psaraiva/d3/pkg/events"

	"github.com/streadway/amqp"
)

type OrderUpdatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderUpdatedHandler(rabbitMQChannel *amqp.Channel) *OrderUpdatedHandler {
	return &OrderUpdatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderUpdatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order updated: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
