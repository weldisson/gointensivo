package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/weldisson/gointensivo/internal/order/entity"
)

func Publish(ch *amqp.Channel, order entity.Order) error {
	//transforma a order em json
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	// publica no rabbitmq na exchange escolhida
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func GenerateOrders() entity.Order {
	return entity.Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
		Tax:   rand.Float64() * 10,
	}
}
func main() {
	// abrir conexão
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	// depois q tudo executar fecha a conexão
	defer conn.Close()
	// abre canal
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 100; i++ {
		Publish(ch, GenerateOrders())
		time.Sleep(300 * time.Millisecond)
	}
}
