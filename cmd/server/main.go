package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	address := "amqp://guest:guest@localhost:5672"
	conn, err := amqp.Dial(address)
	if err != nil {
		fmt.Printf("could not create connection, error: %e\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connection successful!")

	rabbitChan, err := conn.Channel()
	if err != nil {
		fmt.Printf("could not create channel, error: %e\n", err)
		return
	}

	err = pubsub.PublishJSON(
		rabbitChan,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		fmt.Printf("tried to publish json, error: %e\n", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Println("Program closing...")
}
