package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/weldisson/gointensivo/internal/order/infra/database"
	"github.com/weldisson/gointensivo/internal/order/usecase"
	"github.com/weldisson/gointensivo/pkg/rabbitmq"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	// abre um canal
	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}
	// depois que executar fecha a conexão
	defer ch.Close()
	// abrindo um canal Go
	out := make(chan amqp.Delivery) // channel
	// gera uma tread consumindo a mensagem na thread 2 jogando para o channel acima.
	go rabbitmq.Consume(ch, out) // T2
	//imprime todas as mensagens enviadas
	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		// cada mensagem recebida no Body vai ser inserida no InputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		// salva no banco a mensagem
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		// retira a mensagem da fila
		msg.Ack(false)
		//printa a mensagem
		fmt.Println(outputDTO)
	}
}
