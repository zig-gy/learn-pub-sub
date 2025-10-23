package main

import (
	"fmt"
	"os"
	"os/signal"

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
		fmt.Printf("could not create channel, error: %e", err)
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Println("Program closing...")
}
