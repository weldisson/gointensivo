package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChanel() (*amqp.Channel, error) {
	// abre conexão
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		// faz o programa travar caso haja erro de conexão
		panic(err)
	}

	// cria um canal de comunicação amqp
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

// funcao q consome mensagens, joga para o channel do GO,
func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"orders",      // fila
		"go-consumer", // aplicação
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	// loop infinito que lê todas as mensagens e joga para o "out"
	for msg := range msgs {
		out <- msg
	}

	return nil
}
