package main

import (
	"fmt"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
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

	gamelogic.PrintServerHelp()
	continue_loop := true
	for continue_loop {
		inputs := gamelogic.GetInput()
		switch inputs[0] {
		case "pause":
			fmt.Println("Sending pause message")

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

		case "resume":
			fmt.Println("Sending resume message")

			err = pubsub.PublishJSON(
				rabbitChan,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{
					IsPaused: false,
				},
			)
			if err != nil {
				fmt.Printf("tried to publish json, error: %e\n", err)
			}

		case "quit":
			fmt.Println("Exiting program")
			continue_loop = false

		default:
			fmt.Println("Could not understand message")
		}

	}

}
