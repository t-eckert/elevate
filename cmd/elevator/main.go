package main

import (
	"context"
	"log"

	"github.com/t-eckert/elevate/internal/elevator"
)

func main() {
	log.Println("Starting elevator")
	ctx := context.Background()

	elevator := elevator.NewElevator()
	log.Printf("Elevator created with ID: %d", elevator.ID)

	err := elevator.Serve(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
