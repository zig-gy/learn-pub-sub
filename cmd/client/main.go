package main

import (
	"fmt"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
)

func main() {
	fmt.Println("Starting Peril client...")

	name, err := gamelogic.ClientWelcome()
	if err != nil {
		fmt.Printf("error reading username, error: %e", err)
		return
	}

}
