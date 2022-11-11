package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	quantityWorkers := 150
	// load balance
	//quando é executado tem a quantidade de workers recebendo as mensagens do rabbitMQ de forma paralela
	for i := 1; i < quantityWorkers; i++ {
		go worker(out, &uc, i)
	}
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUC := usecase.GetTotalUseCase{OrderRepository: repository}
		total, err := getTotalUC.Execute()
		if err != nil {
			// retorna error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(total)
	})

	http.ListenAndServe(":8080", nil) // chama o server HTTP, cria um thread.
}

// recebe a mensagem que vai ser entregue, use case e id
func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerID int) {
	// imprime todas as mensagens enviadas
	for msg := range deliveryMessage {
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
		fmt.Printf("Worker %d has processsed order %s\n", workerID, outputDTO.ID)
		// aguarda 500 ms
		time.Sleep(500 * time.Microsecond)
	}
}
