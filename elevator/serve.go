package elevator

import (
	"context"
	"time"
)

// Serve runs the elevator continuously, moving the elevator up and down,
// picking up passengers, and dropping them off.
func Serve(ctx context.Context, e *Elevator) {
	// Move the elevator
	go func() {
		for {
			e.Navigate()
			e.Move()

			select {
			case <-time.After(10 * time.Millisecond):
			case <-ctx.Done():
				return
			}
		}
	}()

	// Serve the passengers
	go func() {
		for {
			path := e.Path
			if len(path) == 0 {
				continue
			}

			if e.AtFloor(path[0]) {
				e.PickupAndDropoff()
			}

			select {
			case <-time.After(time.Millisecond):
			case <-ctx.Done():
				return
			}
		}
	}()
}
