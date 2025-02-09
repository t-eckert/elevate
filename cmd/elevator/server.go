package main

import (
	"context"
	"log"
	"time"

	"github.com/t-eckert/elevate/internal/elevator"
)

type Server struct{}

func (s *Server) Serve(ctx context.Context) error {
	log.Println("Starting elevator")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	elevator := elevator.NewElevator(0)
	log.Printf("Elevator created with ID: %d", elevator.ID())

	go runElevator(ctx, elevator)
	go randomlyAddPassengers(ctx, elevator)
	go handleDropoffs(ctx, elevator)
	go logElevatorState(ctx, elevator)

	// Let the system run for a while
	time.Sleep(300 * time.Second)

	// Stop the elevator gracefully
	log.Println("Stopping the elevator")
	cancel()

	// Give time for Goroutines to exit
	time.Sleep(1 * time.Second)

	return nil
}
