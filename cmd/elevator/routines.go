package main

import (
	"context"
	"log"
	"math/rand/v2"
	"time"

	"github.com/t-eckert/elevate/internal/elevator"
	"github.com/t-eckert/elevate/internal/passenger"
)

func runElevator(ctx context.Context, elevator *elevator.Elevator) {
	err := elevator.Serve(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func randomlyAddPassengers(ctx context.Context, elevator *elevator.Elevator) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Duration(rand.IntN(5)+2) * time.Second) // Random interval (2-6 seconds)

			newPassenger := &passenger.Passenger{
				ID:          rand.IntN(1000), // Unique ID
				Origin:      rand.IntN(10),   // Random floor (0-9)
				Destination: rand.IntN(10),   // Random floor (0-9)
			}

			log.Printf("Adding passenger %d going from floor %d to floor %d", newPassenger.ID, newPassenger.Origin, newPassenger.Destination)
			elevator.RequestChan <- newPassenger
		}
	}
}

func handleDropoffs(_ context.Context, elevator *elevator.Elevator) {
	for p := range elevator.DropoffChan {
		log.Printf("Passenger %d dropped off on floor %d", p.ID, p.Destination)
	}
}

func logElevatorState(ctx context.Context, e *elevator.Elevator) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Printf("Elevator:%d;V:%.2f;P:%.2f;On:%d;Rq:%d", e.ID(), e.Velocity()*1000, e.Position(), len(e.Passengers()), len(e.Requests()))
			time.Sleep(time.Second)
		}
	}
}
