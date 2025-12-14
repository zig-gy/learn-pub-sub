package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	name, err := gamelogic.ClientWelcome()
	if err != nil {
		fmt.Printf("error reading username, error: %e", err)
		return
	}

	fmt.Println("Starting Peril client...")

	address := "amqp://guest:guest@localhost:5672"
	conn, err := amqp.Dial(address)
	if err != nil {
		fmt.Printf("could not create connection, error: %e\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connection succesful!")

	_, _, err = pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilDirect,
		fmt.Sprintf("%s.%s", routing.PauseKey, name),
		routing.PauseKey,
		pubsub.Transient,
	)
	if err != nil {
		fmt.Printf("could not declare and bind, error: %e\n", err)
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Println("Program closing...")
}
