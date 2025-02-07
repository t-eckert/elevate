package main

import (
	"context"
	"log"
	"math/rand/v2"
	"time"

	"github.com/t-eckert/elevate/internal/elevator"
	"github.com/t-eckert/elevate/internal/passenger"
)

func main() {
	log.Println("Starting elevator")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	elevator := elevator.NewElevator(0)
	log.Printf("Elevator created with ID: %d", elevator.ID())

	go func() {
		err := elevator.Serve(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Duration(rand.IntN(5)+2) * time.Second) // Random interval (2-6 seconds)

				newPassenger := &passenger.Passenger{
					ID:          rand.IntN(1000), // Unique ID
					Destination: rand.IntN(10),   // Random floor (0-9)
				}

				log.Printf("Adding passenger %d going to floor %d", newPassenger.ID, newPassenger.Destination)
				elevator.AddPassenger(newPassenger)
			}
		}
	}()

	// Let the system run for a while
	time.Sleep(30 * time.Second)

	// Stop the elevator gracefully
	log.Println("Stopping the elevator")
	cancel()

	// Give time for Goroutines to exit
	time.Sleep(1 * time.Second)

}
